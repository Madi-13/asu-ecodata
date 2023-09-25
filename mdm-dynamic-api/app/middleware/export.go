package middleware

//func Monitoring() {}

import (
	"app/domain"
	"app/initializer"
	"bytes"
	"context"
	"encoding/json"

	"log"
	"strconv"

	"net/http"
	"strings"
	"time"

	//"github.com/go-co-op/gocron"
	"github.com/segmentio/kafka-go"
)

const (
	GetHistoryTotalQuery = `select * from mdm_core.func_log_history_form_count()`
	GetHistoryAddQuery   = `select * from mdm_core.func_log_histoty_form_add($1)`
)

func Monitoring() error {

	scenaries := make([]domain.ScenarioEntity, 0)

	con, err := initializer.GetConnection()
	if err != nil {
		log.Println("con -> ", err)
		return err
	}

	rows, err := con.Query(context.Background(),
		"select id, scenario_name, status from mdm_core.export_scenario where status = 'ACTIVE'")
	if err != nil {
		log.Println("rows -> ", err)
		return err
	}

	for rows.Next() {
		scenario := domain.ScenarioEntity{}
		err := rows.Scan(&scenario.ID, &scenario.ScenarioName, &scenario.Status)
		if err != nil {
			log.Println("parse -> ", err)
			return err
		}
		scenaries = append(scenaries, scenario)
	}
	rows.Close()

	for _, scenario := range scenaries {

		prc, err := con.Query(context.Background(),
			`select prc.id, prc.dictionary_id, prc.process_name, prc.process_order, ep.property_name, ep.property_type, ep.path , ep.rq_token , ep.rq_group 
				from mdm_core.export_process prc, mdm_core.export_property ep 
				where ep.id = prc.property_id and  prc.scenario_id = $1
				order by prc.process_order `,
			scenario.ID)
		defer prc.Close()
		if err != nil {
			return err
		}

		processing := make([]domain.ProcessSender, 0)
		for prc.Next() {
			process := domain.ProcessSender{}
			err := prc.Scan(&process.ID,
				&process.DictionaryID,
				&process.ProcessName,
				&process.ProcessOrder,
				&process.PropertyName,
				&process.PropertyType,
				&process.Path,
				&process.PropertyRQToken,
				&process.PropertyRQGroup,
			)
			if err != nil {
				log.Println(err)
				return err
			}
			processing = append(processing, process)
		}
		prc.Close()

		for idx, _ := range processing {

			process := &processing[idx]
			mgs, err := con.Query(context.Background(),
				`select h.id,h.old_value , h.new_value
						  from mdm_data.log_history h
						  where h.object_id = $1 
						  and not exists (select 1 from mdm_core.export_journal j where j.log_id = h.id and j.process_id = $2 )
						  order by h.id limit 200`,
				process.DictionaryID, process.ID)
			if err != nil {
				log.Println(err)
				return err
			}

			messages := make([]domain.Message, 0)

			for mgs.Next() {
				message := domain.Message{}
				err := mgs.Scan(&message.ID, &message.OldValue, &message.NewValue)
				if err != nil {
					log.Println(err)
					return err
				}
				messages = append(messages, message)
			}

			process.Messages = messages
			mgs.Close()

		}

		InsertJournalQuery := `select * from mdm_core.func_export_journal_cud($1, $2)`
		out := domain.GetRecordsOut{}

		for _, prcd := range processing {
			journals := make([]domain.JournalEntity, 0)

			if prcd.PropertyType == "HTTP" {
				journals = append(journals, SendHTTP(prcd.Path, *prcd.PropertyRQToken, prcd)...)
			} else {
				journals = append(journals, SendKAFKA(prcd.Path, *prcd.PropertyRQToken, prcd)...)
			}

			for _, jrn := range journals {

				jrnjs, err := json.Marshal(jrn)
				if err != nil {
					log.Println(err)
					return err
				}

				err = con.QueryRow(
					context.Background(), InsertJournalQuery, string(jrnjs), "SENDER",
				).Scan(&out.DbCode, &out.DbErrorDesc)
				if err != nil {
					log.Println(err)
					return err
				}
			}

		}

	}
	con.Release()

	return nil
}

func SendHTTP(url string, app_type string, prcd domain.ProcessSender) []domain.JournalEntity {

	journals := make([]domain.JournalEntity, 0)
	client := &http.Client{
		Timeout: 2 * time.Second,
	}

	for _, msgs := range prcd.Messages {
		jsonBodySend, err := json.Marshal(msgs)
		if err != nil {
			log.Println(err)
			return nil
		}

		v_meesage_out := ""
		v_status := ""
		body := string(jsonBodySend)

		response, err := client.Post(url, app_type, bytes.NewBuffer([]byte(strings.ReplaceAll(body, "\\", ""))))

		if err != nil {
			log.Println(err)
			v_status = "ERROR"
			v_meesage_out = err.Error()
		} else {
			v_status = "SEND"
			v_meesage_out = response.Status
			response.Body.Close()

		}

		journal := domain.JournalEntity{nil,
			"add",
			prcd.ID,
			msgs.ID,
			v_status,
			"Process name :" + prcd.ProcessName + " Response: " + v_status + " " + v_meesage_out,
			time.Now()}

		journals = append(journals, journal)

	}
	return journals
}

func SendKAFKA(url string, app_token string, prcd domain.ProcessSender) []domain.JournalEntity {

	journals := make([]domain.JournalEntity, 0)

	connection, err := kafka.Dial("tcp", url)
	if err != nil {
		return nil
	}
	defer connection.Close()
	mdmProducerConfig := kafka.TopicConfig{Topic: app_token, NumPartitions: 2, ReplicationFactor: 1}

	err = connection.CreateTopics(mdmProducerConfig)
	if err != nil {
		return nil
	}

	producer := NewProducer(url)

	kfm := make([]kafka.Message, 0)

	for _, msgs := range prcd.Messages {
		jsonBodySend, err := json.Marshal(msgs)
		if err != nil {
			return nil
		}
		body := string(jsonBodySend)

		kfm = append(kfm, kafka.Message{
			Topic:  app_token,
			Offset: 0,
			Key:    []byte(strconv.Itoa(msgs.ID)),
			Value:  []byte(strings.ReplaceAll(body, "\\", "")),
		})

		journal := domain.JournalEntity{nil,
			"add",
			prcd.ID,
			msgs.ID,
			"SEND",
			"Process name :" + prcd.ProcessName,
			time.Now()}

		journals = append(journals, journal)

	}

	err = ProduceBatch(app_token, producer, kfm)
	if err != nil {
		for _, jrn := range journals {
			journal_record := &jrn
			journal_record.Status = "ERROR"
			journal_record.Description = err.Error()
		}
	}

	/*
		var batch_size = 50
		var batch_count = len(kfm) / batch_size

		for idx := 0; idx < batch_count; idx++ {

			kfmb := &kfm[idx*batch_size : batch_size*idx+1]

			err = ProduceBatch(app_token, producer, kfmb)
			if err != nil {
				for _, jrn := range journals {
					journal_record := &jrn
					journal_record.Status = "ERROR"
					journal_record.Description = err.Error()
				}
			}

		}
	*/

	return journals
}

func NewProducer(brokers string) *Producer {
	dialer := &kafka.Dialer{
		Timeout:   10 * time.Second,
		DualStack: true,
	}

	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers:      []string{brokers},
		Dialer:       dialer,
		BatchSize:    25,
		BatchTimeout: 2 * time.Second,
		RequiredAcks: 0,
	})

	return &Producer{
		Writer: writer,
	}
}

type Producer struct {
	Writer *kafka.Writer
	Dialer *kafka.Dialer
}

func Produce(key []byte, value []byte, topic string, producer *Producer) error {
	err := producer.Writer.WriteMessages(context.Background(), kafka.Message{
		Topic:  topic,
		Offset: 0,
		Key:    key,
		Value:  value,
	})

	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func ProduceBatch(topic string, producer *Producer, messages []kafka.Message) error {
	err := producer.Writer.WriteMessages(context.Background(), messages...)

	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func PrepareExport() error {

	log.Println("PrepareExport -> ")
	ImportUsers()
	con, err := initializer.GetConnection()
	if err != nil {
		return err
	}
	defer con.Release()

	total := 0
	err = con.QueryRow(context.Background(), GetHistoryTotalQuery).Scan(&total)
	if err != nil {
		return err
	}

	count := total / 500
	iter := 0
	for count >= iter {
		iter++
		_, err := con.Query(context.Background(), GetHistoryAddQuery, 500)
		if err != nil {
			return err
		}
	}
	return nil

}

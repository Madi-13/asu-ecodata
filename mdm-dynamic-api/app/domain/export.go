package domain

import "time"

type Operation struct {
	OperationType string         `json:"operation_type"`
	RowID         int            `json:"row_id,omitempty"`
	ScenarioID    int            `json:"scenario_id,omitempty"`
	Limit         int            `json:"limit"`
	Offset        int            `json:"offset"`
	Search        string         `json:"search"`
	Property      PropertyEntity `json:"property,omitempty"`
	Scenarion     ScenarioEntity `json:"scenario,omitempty"`
	Process       ProcessEntity  `json:"process,omitempty"`
	Journal       JournalEntity  `json:"journal,omitempty"`
}

type PropertyOut struct {
	Total    int              `json:"total"`
	Property []PropertyEntity `json:"result"`
}

type PropertyEntity struct {
	ID           int     `json:"id"`
	PropertyName string  `json:"property_name"`
	PropertyType string  `json:"property_type"`
	Path         string  `json:"path"`
	RqToken      *string `json:"rq_token"`
	RqGroup      *string `json:"rq_group"`
	PropertyMode string  `json:"property_mode"`
}

type ScenarioOut struct {
	Total    int              `json:"total"`
	Scenario []ScenarioEntity `json:"result"`
}

type ScenarioEntity struct {
	ID           int       `json:"id"`
	ScenarioName string    `json:"scenario_name"`
	Status       string    `json:"status"`
	CreateTS     time.Time `json:"create_date"`
	CreateBy     string    `json:"create_by"`
}

type ProcessOut struct {
	Total   int             `json:"total"`
	Process []ProcessEntity `json:"result"`
}

type ProcessEntity struct {
	ID              int    `json:"id"`
	PropertyID      int    `json:"property_id"`
	ScenarioID      int    `json:"scenario_id"`
	DictionaryID    int    `json:"dictionary_id"`
	ProcessName     string `json:"process_name"`
	ProcessOrder    int    `json:"process_order"`
	ScenarioName    string `json:"scenario_name"`
	PropertyName    string `json:"property_name"`
	Path            string `json:"path"`
	PropertyType    string `json:"property_type"`
	DictionaryName  string `json:"dictionary_name"`
	DictionaryLocal string `json:"dictionary_local"`

	CreateTS time.Time `json:"create_date"`
	CreateBy string    `json:"create_by"`
}

type ProcessSender struct {
	ID              int       `json:"id"`
	PropertyID      int       `json:"property_id"`
	ScenarionID     int       `json:"scenarion_id"`
	DictionaryID    int       `json:"dictionariyes_id"`
	ProcessName     string    `json:"process_name"`
	ProcessOrder    int       `json:"process_order"`
	PropertyName    string    `json:"property_name"`
	PropertyType    string    `json:"property_type"`
	PropertyRQToken *string   `json:"property_rq_token"`
	PropertyRQGroup *string   `json:"property_rq_group"`
	Path            string    `json:"path"`
	Messages        []Message `json:"messages"`
}

type JournalOut struct {
	Total   int             `json:"total"`
	Journal []JournalEntity `json:"result"`
}

type JournalEntity struct {
	ID            *int      `json:"id"`
	OperationType string    `json:"operation_type"`
	ProcessID     int       `json:"process_id"`
	LogID         int       `json:"log_id"`
	Status        string    `json:"status"`
	Description   string    `json:"description"`
	SendDate      time.Time `json:"send_date"`
}

type Message struct {
	ID       int     `json:"id"`
	OldValue *string `json:"old_value"`
	NewValue *string `json:"new_value"`
}

type ImportMonitoringList struct {
	Total            int                `json:"total"`
	ImportMonitoring []ImportMonitoring `json:"import"`
}

type ImportMonitoring struct {
	ID        int        `json:"id"`
	FileName  string     `json:"file_name"`
	DateStart *time.Time `json:"date_start"`
	DateEnd   *time.Time `json:"date_end"`
	States    *string    `json:"states"`
	Status    *string    `json:"status"`
	Messages  *string    `json:"messages"`
}

type RFile struct {
	ObjectId int    `json:"object_id"`
	RowId    int    `json:"row_id"`
	FileId   int    `json:"file_id"`
	Files    []File `json:"files"`
	UUID     string
}

type File struct {
	FileName string `json:"file_name"`
	FileBody string `json:"file_body"`
}

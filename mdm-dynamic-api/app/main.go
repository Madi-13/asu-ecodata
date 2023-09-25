package main

import (
	"app/controllers"
	"app/domain"
	"app/initializer"
	"app/middleware"
	"fmt"
	"log"

	"net/http"
	"time"

	"github.com/go-co-op/gocron"
	//"github.com/gofiber/fiber/v2"
	//"github.com/gofiber/fiber/v2/middleware/logger"
	//"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gorilla/mux"
)

var (
	config *domain.Config
	err    error
)

func init() {
	config, err = initializer.LoadConfig("./config")
	if err != nil {
		log.Fatal("🚀 Could not load environment variables", err)
	}

	DbUrl := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		config.DBUserName, config.DBUserPassword, config.DBHost,
		config.DBPort, config.DBName,
	)
	err = initializer.ConnectDB(&DbUrl)
	if err != nil {
		log.Println("initializers.ConnectDB error:", err)
	}

}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.RequestURI)
		next.ServeHTTP(w, r)
	})
}

func main() {
	log.Println("Starting dynamic MDM")

	defer initializer.PoolClose()
	go ScheduleMonitoring()

	route := mux.NewRouter()

	route.Use(loggingMiddleware)
	route.HandleFunc("/healthz", controllers.Healthzv1)
	mapp := route.PathPrefix("/api").Subrouter()
	table := mapp.PathPrefix("/objects").Subrouter()
	table.HandleFunc("/tables", controllers.GetTablesV2)
	table.HandleFunc("/tables_menu", controllers.GetTablesByTypeV2)
	table.HandleFunc("/tables/struct", controllers.GetTableStructV2)
	table.HandleFunc("/tables/file_upload", controllers.FileUploaderV2)
	table.HandleFunc("/tables/file_list", controllers.FilesListV2)
	table.HandleFunc("/tables/delete_file", controllers.DeleteFilesV2)

	table.HandleFunc("/table", controllers.CreateTableV2)
	table.HandleFunc("/table/alter", controllers.AlterTableV2)
	table.HandleFunc("/table/drop", controllers.DropTableV2)
	table.HandleFunc("/types", controllers.GetTypesV2)

	minfo := mapp.PathPrefix("/info").Subrouter()
	minfo.HandleFunc("/date", controllers.InfoByDateV2)
	minfo.HandleFunc("/main", controllers.InfoMainV2)
	minfo.HandleFunc("/filter", controllers.FilterListV2)

	mrecords := mapp.PathPrefix("/records").Subrouter()
	mrecords.HandleFunc("", controllers.GetRecordsV2)
	mrecords.HandleFunc("/set_status", controllers.SetRecordStatusV2)
	mrecords.HandleFunc("/tree", controllers.GetRecordsTreeV2)
	mrecords.HandleFunc("/for_edit", controllers.GetRecordForEditV2)
	mrecords.HandleFunc("/for_create", controllers.GetRecordForCreateV2)
	mrecords.HandleFunc("/create", controllers.GetCUDOperationV2)
	mrecords.HandleFunc("/update", controllers.GetCUDOperationV2)
	mrecords.HandleFunc("/delete", controllers.GetCUDOperationV2)
	mrecords.HandleFunc("/popup", controllers.GetRecordForPopupV2)
	mrecords.HandleFunc("/history", controllers.GetLogHistoryV2)

	mexport := mapp.PathPrefix("/export").Subrouter()
	mexport.HandleFunc("/scenario", controllers.GetScenarioRecordsV2)
	mexport.HandleFunc("/scenario/dml", controllers.GetCUDScenarioV2)
	mexport.HandleFunc("/property", controllers.GetPropertyRecordsV2)
	mexport.HandleFunc("/property/dml", controllers.GetCUDPropertyV2)
	mexport.HandleFunc("/process", controllers.GetProcessRecordsV2)
	mexport.HandleFunc("/process/dml", controllers.GetCUDProcessV2)
	mexport.HandleFunc("/journal", controllers.GetJournalRecordsV2)
	mexport.HandleFunc("/journal/dml", controllers.GetCUDJournalV2)
	mexport.HandleFunc("/replicate", controllers.GetReplicateRecordV2)
	mexport.HandleFunc("/values", controllers.GetExportTableValuesV2)

	mattribute := mapp.PathPrefix("/attribute").Subrouter()
	mattribute.HandleFunc("", controllers.GetAttributeRecordsV2)
	mattribute.HandleFunc("/values", controllers.GetAttributeValuesV2)
	mattribute.HandleFunc("/values/create", controllers.GetFastCreateAttributeValuesV2)
	mattribute.HandleFunc("/dml", controllers.GetCUDAttributeV2)
	mattribute.HandleFunc("/values/dml", controllers.GetCUDAttributeValuesV2)

	mimport := mapp.PathPrefix("/import").Subrouter()
	mimport.HandleFunc("", controllers.ImportTableV2)
	mimport.HandleFunc("/monitoring", controllers.GetImportMonitoringV2)
	mimport.HandleFunc("/monitoring/ends", controllers.GetImportMonitoringEndingV2)
	mimport.HandleFunc("/tamplate", controllers.ImportTamplateTableV2)

	madmin := mapp.PathPrefix("/admin").Subrouter()
	madmin.HandleFunc("/user/find", controllers.FindUsersV2)

	srv := &http.Server{
		Handler: route,
		Addr:    config.ServerIp + ":" + config.ServerPort,

		WriteTimeout: 45 * time.Second,
		ReadTimeout:  45 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())

	/*app := fiber.New()
	app.Use(logger.New())

	app.Get("/metrics", monitor.New(monitor.Config{Title: "MyService Metrics Page"}))
	app.Get("/healthz", controllers.Healthz)

	api := app.Group("/api")
	objects := api.Group("/objects")

	objects.Post("/tables", middleware.CheckToken,
		controllers.GetTables)
	objects.Post("/tables_menu", //middleware.CheckToken,
		controllers.GetTablesByType)
	objects.Post("/tables/struct", // middleware.CheckToken,
		controllers.GetTableStruct)

	objects.Post("/tables/file_upload", middleware.CheckToken,
		controllers.FileUploader)
	objects.Post("/tables/file_list", middleware.CheckToken,
		controllers.FilesList)
	objects.Post("/tables/delete_file", middleware.CheckToken,
		controllers.DeleteFiles)

	objects.Post("/table", middleware.CheckToken,
		controllers.CreateTable)
	objects.Post("/table/alter", middleware.CheckToken,
		controllers.AlterTable)
	objects.Post("/table/drop", middleware.CheckToken,
		controllers.DropTable)

	objects.Get("/types", //middleware.CheckToken,
		controllers.GetTypes)

	info := api.Group("/info")
	info.Post("/date", controllers.InfoByDate)
	info.Post("/main", controllers.InfoMain)
	info.Post("/filter", controllers.FilterList)

	records := api.Group("/records")
	records.Post(
		"/", //middleware.CheckToken, //middleware.SetRead, middleware.CheckPermission,
		controllers.GetRecords,
	)

	records.Post(
		"/set_status", middleware.CheckToken, //middleware.SetRead, middleware.CheckPermission,
		controllers.SetRecordStatus,
	)

	records.Post(
		"/tree/", //middleware.CheckToken, // middleware.SetRead, middleware.CheckPermission,
		controllers.GetRecordsTree,
	)

	records.Post(
		"/for_edit", //middleware.CheckToken, //middleware.SetRead, middleware.CheckPermission,
		controllers.GetRecordForEdit,
	)

	records.Post(
		"/for_create", //middleware.CheckToken, // middleware.SetRead, middleware.CheckPermission,
		controllers.GetRecordForCreate,
	)
	records.Post(
		"/create/", middleware.CheckToken, // middleware.SetWrite, middleware.CheckPermission,
		controllers.GetCUDOperation,
	)
	records.Post(
		"/update/", middleware.CheckToken, // middleware.SetWrite, middleware.CheckPermission,
		controllers.GetCUDOperation,
	)
	records.Post(
		"/delete/", middleware.CheckToken, // middleware.SetDelete, middleware.CheckPermission,
		controllers.GetCUDOperation,
	)

	records.Post(
		"/popup/", //middleware.CheckToken, //middleware.SetRead, middleware.CheckPermission,
		controllers.GetRecordForPopup,
	)
	records.Post(
		"/history/", //middleware.CheckToken, //middleware.SetRead, middleware.CheckPermission,
		controllers.GetLogHistory,
	)

	export := api.Group("/export")
	export.Post("/scenario", // middleware.CheckToken,
		controllers.GetScenarioRecords)
	export.Post("/scenario/dml", // middleware.CheckToken,
		controllers.GetCUDScenario)

	export.Post("/property", // middleware.CheckToken,
		controllers.GetPropertyRecords)
	export.Post("/property/dml", // middleware.CheckToken,
		controllers.GetCUDProperty)

	export.Post("/process", // middleware.CheckToken,
		controllers.GetProcessRecords)
	export.Post("/process/dml", // middleware.CheckToken,
		controllers.GetCUDProcess)

	export.Post("/journal", // middleware.CheckToken,
		controllers.GetJournalRecords)
	//export.Post("/journal/dml", // middleware.CheckToken,
	//	controllers.GetCUDJournal)

	export.Post("/replicate", // middleware.CheckToken,
		controllers.GetReplicateRecord)
	export.Post("/values", // middleware.CheckToken,
		controllers.GetExportTableValues)

	attribute := api.Group("/attribute")

	attribute.Post("/", //middleware.CheckToken,
		controllers.GetAttributeRecords)

	attribute.Post("/values", //middleware.CheckToken,
		controllers.GetAttributeValues)

	attribute.Post("/values/create", middleware.CheckToken,
		controllers.GetFastCreateAttributeValues)

	attribute.Post("/dml", middleware.CheckToken,
		controllers.GetCUDAttribute)

	attribute.Post("/values/dml", middleware.CheckToken,
		controllers.GetCUDAttributeValues)

	vimport := api.Group("/import")
	vimport.Post("/", middleware.CheckToken,
		controllers.ImportTable)
	vimport.Post("/monitoring", // middleware.CheckToken,
		controllers.GetImportMonitoring)

	vimport.Post("/monitoring/ends", // middleware.CheckToken,
		controllers.GetImportMonitoringEnding)
	vimport.Post("/tamplate", // middleware.CheckToken,
		controllers.ImportTamplateTable)

	admin := api.Group("/admin")

	admin.Post("/user/find", //middleware.CheckToken,
		controllers.FindUsers)



	err := app.Listen(":" + config.ServerPort)
	if err != nil {
		panic(err)
	}
	*/
}

func ScheduleMonitoring() {

	sch := gocron.NewScheduler(time.UTC)
	log.Println("PREPARE SENDING SCHEDULER ")
	SendFunc := middleware.Monitoring
	_, err := sch.Every(60).Second().Do(SendFunc)
	if err != nil {
		log.Println("Error sending ", err.Error())
	}
	PrepareFunc := middleware.PrepareExport
	_, err = sch.Every(600).Second().Do(PrepareFunc)
	if err != nil {
		log.Println("Error prepare ", err.Error())
	}

	/*
		UsersFunc := middleware.ImportUsers
		_, err = sch.Every(6000).Second().Do(UsersFunc)
		if err != nil {
			log.Println("Error users ", err.Error())
		}
	*/

	sch.StartAsync()

}

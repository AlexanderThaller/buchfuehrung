package main

import (
	"flag"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
)

import (
	data "github.com/AlexanderThaller/buchfuehrung"
	"github.com/AlexanderThaller/httphelper"
	log "github.com/Sirupsen/logrus"
	"github.com/jinzhu/gorm"
	"github.com/juju/errgo"
	"github.com/julienschmidt/httprouter"
	_ "github.com/lib/pq"
	"github.com/prometheus/client_golang/prometheus"
)

var (
	BuildTime    string
	BuildVersion string

	FlagBindingAPI     string
	FlagBindingMetrics string

	FlagDatabaseType             string
	FlagDatabaseConnectionString string

	FlagLogLevel string
	FlagLogFile  string

	FlagPIDFile string

	Database gorm.DB
)

func init() {
	// Binding
	flag.StringVar(&FlagBindingAPI, "binding.files", ":5523",
		"the network binding of the api")
	flag.StringVar(&FlagBindingMetrics, "binding.metrics", ":5524",
		"the network binding of the metrics")

	// Database
	flag.StringVar(&FlagDatabaseType, "database.type", "postgres",
		"the network binding of the api")
	flag.StringVar(&FlagDatabaseConnectionString, "database.connection", "user=postgres password=postgres dbname=buchfuehrung",
		"the network binding of the metrics")

	//Log
	flag.StringVar(&FlagLogLevel, "log.level", "info",
		"the loglevel of the application (debug, info, warning, error")
	flag.StringVar(&FlagLogFile, "log.file", "",
		"the path to the logfile. if empty logs will go to stdout and not a logfile")

	// PID
	flag.StringVar(&FlagPIDFile, "pid.file", "",
		"write the pid of the current process to the given file")

	flag.Parse()

	level, err := log.ParseLevel(FlagLogLevel)
	if err != nil {
		log.Fatal(errgo.Notef(err, "can not parse loglevel from flag"))
	}
	log.SetLevel(level)

	if FlagPIDFile != "" {
		pid := strconv.Itoa(os.Getppid())
		log.Debug("PID: ", pid)

		err := ioutil.WriteFile(FlagPIDFile, []byte(pid), 0644)
		if err != nil {
			log.Fatal(errgo.Notef(err, "can not write pid to file"))
		}
	}
}

func main() {
	if FlagLogFile != "" {
		logfile, err := os.OpenFile(FlagLogFile, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0644)
		if err != nil {
			log.Fatal(errgo.Notef(err, "can not open logfile for writing"))
		}
		defer logfile.Close()

		log.SetOutput(logfile)
	}

	log.Info("Starting buchführung api v", BuildVersion, " +", BuildTime)

	err := initDatabase()
	if err != nil {
		log.Fatal(errgo.Notef(err, "can not initialize database"))
	}

	go func() {
		router := httprouter.New()

		// Router handler
		router.MethodNotAllowed = httphelper.HandlerLoggerHTTP(httphelper.PageRouterMethodNotAllowed)
		router.NotFound = httphelper.HandlerLoggerHTTP(httphelper.PageRouterNotFound)

		// Root and Favicon
		router.GET("/", httphelper.HandlerLoggerRouter(pageRoot))
		router.GET("/favicon.ico", httphelper.HandlerLoggerRouter(httphelper.PageMinimalFavicon))

		// API v0
		// Accounts
		router.POST("/api/v0/account/add", httphelper.HandlerLoggerRouter(pageAPIV0AccountAdd))
		router.GET("/api/v0/account/get/byid/:id", httphelper.HandlerLoggerRouter(pageAPIV0AccountGetByID))
		router.GET("/api/v0/account/get/byname/:name", httphelper.HandlerLoggerRouter(pageAPIV0AccountGetByName))
		router.GET("/api/v0/account/list", httphelper.HandlerLoggerRouter(pageAPIV0AccountList))

		// Transactions
		router.POST("/api/v0/transaction/add", httphelper.HandlerLoggerRouter(pageAPIV0TransactionAdd))
		router.GET("/api/v0/transaction/get/byid/:id", httphelper.HandlerLoggerRouter(pageAPIV0TransactionGetByID))
		router.GET("/api/v0/transaction/list", httphelper.HandlerLoggerRouter(pageAPIV0TransactionList))

		log.Info("Start serving api on ", FlagBindingAPI)
		log.Fatal(http.ListenAndServe(FlagBindingAPI, router))
	}()

	go func() {
		if FlagBindingMetrics != "" {
			log.Info("Starting Metrics", FlagBindingMetrics)
			http.Handle("/metrics", prometheus.Handler())
			http.ListenAndServe(FlagBindingMetrics, nil)
		}
	}()

	log.Debug("Waiting for interrupt signal")
	httphelper.WaitForStopSignal()
	log.Info("Stopping")
}

func initDatabase() error {
	var err error

	Database, err = gorm.Open(FlagDatabaseType, FlagDatabaseConnectionString)
	if err != nil {
		return errgo.Notef(err, "can not open database")
	}

	log.Debug("Creating table Accounts")
	Database.AutoMigrate(&data.Account{})
	Database.AutoMigrate(&data.Budget{})
	Database.AutoMigrate(&data.BudgetGroup{})
	Database.AutoMigrate(&data.BudgetCategory{})
	Database.AutoMigrate(&data.Category{})
	Database.AutoMigrate(&data.Payee{})
	Database.AutoMigrate(&data.Transaction{})

	return nil
}

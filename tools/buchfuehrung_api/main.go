package main

import (
	"flag"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
)

import (
	"github.com/AlexanderThaller/httphelper"
	log "github.com/Sirupsen/logrus"
	"github.com/juju/errgo"
	"github.com/julienschmidt/httprouter"
	"github.com/prometheus/client_golang/prometheus"
)

var (
	BuildTime    string
	BuildVersion string

	FlagBindingAPI     string
	FlagBindingMetrics string

	FlagLogLevel string
	FlagLogFile  string

	FlagPIDFile string
)

func init() {
	flag.StringVar(&FlagBindingAPI, "binding.files", ":5523",
		"the network binding of the api")
	flag.StringVar(&FlagBindingMetrics, "binding.metrics", ":5524",
		"the network binding of the metrics")

	flag.StringVar(&FlagLogLevel, "log.level", "info",
		"the loglevel of the application (debug, info, warning, error")
	flag.StringVar(&FlagLogFile, "log.file", "",
		"the path to the logfile. if empty logs will go to stdout and not a logfile")

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

	go func() {
		router := httprouter.New()

		// Router handler
		router.MethodNotAllowed = httphelper.HandlerLoggerHTTP(httphelper.PageRouterMethodNotAllowed)
		router.NotFound = httphelper.HandlerLoggerHTTP(httphelper.PageRouterNotFound)

		// Root and Favicon
		router.GET("/", httphelper.HandlerLoggerRouter(pageRoot))
		router.GET("/favicon.ico", httphelper.HandlerLoggerRouter(httphelper.PageMinimalFavicon))

		// API v0
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

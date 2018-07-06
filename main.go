package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gorilla/mux"
	"go.uber.org/zap"

	"go-elastic-content-service/config"
	"go-elastic-content-service/content"

	logger "github.com/ricardo-ch/go-logger"
	tracing "github.com/ricardo-ch/go-tracing"
)

const appName = "go-elastic-content-service"

func init() {
	// initialization (optional)
	logger.InitLogger(false)
}

func main() {

	//Zipkin Connection
	tracing.SetGlobalTracer(appName, config.SvcTracingZipkin)
	defer tracing.FlushCollector()

	// Errors channel
	errc := make(chan error)

	// Interrupt handler.
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errc <- fmt.Errorf("%s", <-c)
	}()

	// content endpoint
	contentService := content.NewService(content.NewRepository())
	contentService = content.NewTracing(contentService)
	contentHandler := content.NewHandler(contentService)

	go func() {

		httpAddr := ":" + config.AppPort
		router := mux.NewRouter()

		// index endpoint
		router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			fmt.Fprintln(w, fmt.Sprintf("Welcome to the %s API!", appName))
		})

		// healthz endpoint
		router.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			w.WriteHeader(http.StatusOK)
		})

		router.Handle("/content/", tracing.HTTPMiddleware("content-handler", http.HandlerFunc(contentHandler.Get))).Methods("GET")

		router.Handle("/content/{id}", tracing.HTTPMiddleware("content-handler", http.HandlerFunc(contentHandler.Put))).Methods("PUT")
		router.Handle("/content/{id}", tracing.HTTPMiddleware("content-handler", http.HandlerFunc(contentHandler.Delete))).Methods("DELETE")

		httpServer := &http.Server{
			Addr:    httpAddr,
			Handler: router,
		}

		logger.Info(fmt.Sprintf("The microservice %s is started on port %s", appName, config.AppPort), zap.String("port", config.AppPort))
		errc <- httpServer.ListenAndServe()

	}()

	logger.Error("exit", zap.Error(<-errc))
}

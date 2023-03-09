package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"ramazon/api"
	"ramazon/ci"
	"ramazon/configs"
	"syscall"

	// "mahalla/pkg/logger"
	"ramazon/pkg/middleware"
	"ramazon/service"

	"github.com/blendle/zapdriver"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

func main() {
	conf := configs.Load()
	if err := conf.Validate(); err != nil {
		panic(err)
	}

	logger, err := zapdriver.NewDevelopment() // with `development` set to `true`
	if err != nil {
		panic(err)
	}

	ci.MigrationsUp()
	ramazonService := service.NewRamazonService(logger)

	root := mux.NewRouter()

	root.Use(middleware.PanicRecovery)
	root.Use(middleware.Logging)

	api.Init(root, ramazonService, logger)

	log.Println("main: Project is started on the port: ", conf.HTTPPort)

	errChan := make(chan error, 1)
	osSignals := make(chan os.Signal, 1)
	signal.Notify(osSignals, os.Interrupt, syscall.SIGTERM)

	httpServer := http.Server{
		Addr:    conf.HTTPPort,
		Handler: root,
	}

	// http server
	go func() {
		errChan <- httpServer.ListenAndServe()
	}()

	// Blocking main and waiting for shutdown.
	select {
	case err := <-errChan:
		logger.Fatal("error: ", zap.Error(err))

	case <-osSignals:
		logger.Info("main : recieved os signal, shutting down")
		_ = httpServer.Shutdown(context.Background())
		return
	}
}

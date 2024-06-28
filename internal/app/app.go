package app

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"card-detect-demo/internal/controller/router"
	"card-detect-demo/internal/service"
	"card-detect-demo/internal/service/onnx"
)

type app struct {
	config *Config
}

func NewApp(config *Config) *app {
	return &app{
		config: config,
	}
}

func (a *app) Run() error {

	// services
	onnxRecognizer, err := onnx.NewService(a.config.Onnx.PathRuntime, a.config.Onnx.PathModel, a.config.IsLogTime)
	if err != nil {
		log.Println(err)
		return err
	}
	detectService := service.NewDetector(onnxRecognizer, a.config.StorageFolder, a.config.IsLogTime)

	// handlers
	h := router.NewRouter(detectService, a.config.StorageFolder, a.config.Name, a.config.Version)

	// start server
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", a.config.Port),
		Handler: h,
	}
	go func() {
		log.Println("Starting app:", a.config.Name, a.config.Version)
		log.Println("Listening on port", a.config.Port)
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatalf("ListenAndServe(): %v", err)
		}
	}()

	// Set up a channel to listen for an interrupt or terminate signal from the OS.
	// Use a buffered channel to prevent missing signals.
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	// Block until a signal is received.
	<-stop

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Doesn't block if no connections, but will otherwise wait until the timeout deadline.
	log.Printf("Shutting down server...")
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server Shutdown: %v", err)
	}

	log.Printf("Server gracefully stopped")
	return nil
}

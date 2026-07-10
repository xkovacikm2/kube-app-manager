package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/xkovacikm2/kube-app-manager/internal/httpapi"
	"github.com/xkovacikm2/kube-app-manager/internal/kube"
)

const (
	defaultPort     = "8080"
	shutdownTimeout = 10 * time.Second
)

func main() {
	port := strings.TrimSpace(os.Getenv("PORT"))
	if port == "" {
		port = defaultPort
	}

	applicationSource, err := kube.NewWebEndpointSource()
	if err != nil {
		log.Fatalf("failed to initialize kubernetes source: %v", err)
	}

	server := &http.Server{
		Addr:              ":" + port,
		Handler:           httpapi.NewHandler(applicationSource),
		ReadHeaderTimeout: 5 * time.Second,
	}

	serverErrorChannel := make(chan error, 1)
	go func() {
		log.Printf("listening on %s", server.Addr)
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			serverErrorChannel <- err
		}
	}()

	signalContext, stopSignals := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	defer stopSignals()

	select {
	case <-signalContext.Done():
		log.Print("shutdown signal received")
	case err := <-serverErrorChannel:
		log.Fatalf("http server failed: %v", err)
	}

	shutdownContext, cancelShutdown := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancelShutdown()

	if err := server.Shutdown(shutdownContext); err != nil {
		log.Fatalf("server shutdown failed: %v", err)
	}
}

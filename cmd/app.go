package main

import (
	"boost/pkg/config"
	"boost/pkg/utils"
	"boost/server"
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	cfg := config.GetConfig()
	logger := utils.GetLogger(cfg.Log.Path, cfg.Log.Filename, cfg.IsDebug)

	srv := server.NewServer(ctx, logger, cfg)
	srv.Serve()

	// When a shutdown signal (like Ctrl+C) is caught, you initiate a graceful shutdown using srv.Shutdown(ctx),
	// giving active connections time to complete before the server shuts down.
	// Graceful shutdown is handled within the main function, allowing ongoing requests to finish processing
	// before the server shuts down, if req not completed in 5 seconds the server will force shutdown.
	// If the expected requests finishes before 5 seconds, the server is shut down immediately.
	// New Requests will not be accepted for 5 seconds.
	// Wait for an interrupt signal to gracefully shutdown the server
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	<-c // Wait for a signal

	// Graceful shutdown
	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdownCancel()
	srv.GracefulShutdown(shutdownCtx)
	logger.Println("Server exiting")
}

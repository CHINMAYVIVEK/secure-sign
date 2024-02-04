package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"secure-sign/config"
	"secure-sign/helper"
	"secure-sign/middleware"
	"syscall"
	"time"
)

// StartServer initializes and starts the HTTP server with graceful shutdown.
func StartServer() {

	// Initialize the router and apply middleware
	r := middleware.ApplyMiddleware(NewRouter(), middleware.LoggerMiddleware)

	// Set up server configurations
	serverAddr := fmt.Sprintf(":%s", config.Cfg.Server.Port)

	// Create a new HTTP server with the router
	server := &http.Server{
		Addr:         serverAddr,
		Handler:      r,
		IdleTimeout:  120 * time.Second, // Adjust as needed
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	// Set up signal handling for graceful shutdown
	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, os.Interrupt, syscall.SIGTERM)

	// Start the server in a goroutine
	go func() {
		helper.SugarObj.Info("Starting server...")
		log.Printf("Server is running on http://localhost%s\n", serverAddr)
		helper.SugarObj.Info("Server is running on http://localhost%s", serverAddr)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Println("Error starting server:", err)
		}
	}()

	// Wait for interrupt signal for graceful shutdown
	<-stopChan
	helper.SugarObj.Info("Shutting down server gracefully...")
	log.Println("Shutting down server gracefully...")

	// Create a context with timeout for shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Shutdown the server gracefully
	if err := server.Shutdown(ctx); err != nil {
		helper.SugarObj.Error("Error shutting down server:", err)

	}
	helper.SugarObj.Info("Server gracefully shut down.")
	log.Println("Server gracefully shut down.")
}

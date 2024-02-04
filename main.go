package main

import (
	"runtime"

	"secure-sign/config"
	"secure-sign/server"
)

func main() {
	// Set GOMAXPROCS to utilize available CPU cores
	runtime.GOMAXPROCS(runtime.NumCPU())

	// Initialize the database connection
	config.ConnectPostgreSQL()

	// Start the server
	server.StartServer()
}

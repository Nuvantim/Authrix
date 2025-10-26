package main

import (
	"api/config"
	"api/database"
	rds "api/cache"
	"api/internal/server/http"
	"api/pkgs/guards"

	"log"
)

func main() {
	// Check Environment
	if err := config.CheckEnv(); err != nil {
		log.Fatal(err)
	}
	// Generate RSA
	guard.GenRSA()
	// Get RSA
	guard.CheckRSA()

	// Start Server
	app := http.ServerGo()

	// Get Server Config
	serverConfig, err := config.GetServerConfig()
	if err != nil {
		log.Fatal(err)
	}

	done := make(chan bool, 1)

	// Start server in goroutine
	go func() {
		if err := app.Listen(":" + serverConfig.Port); err != nil {
			log.Printf("Server stopped: %v", err)
			done <- true
		}
	}()

	config.GracefulShutdown(app, done)

	<-done
	// Close Connection redis
	rds.RedisClose()
	// Close Connection database
	database.CloseDB()
}

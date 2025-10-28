package main

import (
	"api/config"
	"api/database"
	"api/internal/server/http"
	"api/pkgs/guards"
	rds "api/redis"

	"log"
	"os"
	"os/exec"
	"runtime"
)

func ClearScreen() {
	switch runtime.GOOS {
	case "windows":
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Run()
	case "linux", "darwin":
		cmd := exec.Command("/usr/bin/clear")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Run()
	default:
		log.Fatal("OS not detected")
	}
}

func main() {
	// Check Environment
	if err := config.CheckEnv(); err != nil {
		log.Fatal(err)
	}
	// Generate RSA
	guard.GenRSA()
	// Get RSA
	guard.CheckRSA()

	// clear screen
	ClearScreen()

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

package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"ispmyadmin/server"
	"ispmyadmin/utils"
)

func main() {
	showSpirit()

	// Initialize logger
	server.InitLogger()

	// Initialize database
	err := server.InitDatabase()
	if err != nil {
		log.Fatal("Database error:", err)
	}

	// Start RADIUS server in a goroutine
	go func() {
		err := server.StartRadiusServer()
		if err != nil {
			log.Fatal("RADIUS server failed:", err)
		}
	}()

	// Graceful shutdown handling
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-stop
		log.Println("Shutting down gracefully...")
		server.CloseDatabase() // Cleanup database
		os.Exit(0)
	}()

	// Start Admin CLI
	utils.StartAdminCLI()
}

// showSpirit displays the welcome banner
func showSpirit() {
	fmt.Println(`
.___  ___________________                _____       .___      .__        
|   |/   _____/\______   \_____ ___.__. /  _  \    __| _/_____ |__| ____  
|   |\_____  \  |     ___/     <   |  |/  /_\  \  / __ |/     \|  |/    \ 
|   |/        \ |    |  |  Y Y  \___  /    |    \/ /_/ |  Y Y  \  |   |  \
|___/_______  / |____|  |__|_|  / ____\____|__  /\____ |__|_|  /__|___|  /
            \/                \/\/            \/      \/     \/        \/ 
   WELCOME TO ISPmyAdmin — Simplicity is Power — DT7 LEGACY STARTS HERE
`)
}

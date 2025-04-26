package main

import (
	"inine-track/pkg/config"
	"inine-track/pkg/database"
	"inine-track/pkg/routes"
	"log"
)

func main() {

	err := database.ConnectDB()

	if err != nil {
		log.Fatal("erro ao conectar com o banco!", err.Error())
	}
	if err := config.LoadSendEmail(); err != nil {
		log.Fatalf("Error loading configuration: %v", err)
	}

	routes.HandlleRequest()

}

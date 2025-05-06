package main

import (
	"inine-track/pkg/config"
	"inine-track/pkg/database"
	"inine-track/pkg/routes"
	"log"
)

// @title API Inine-Track
// @version 1.0
// @description Esta é uma API feita para análise de dos projetos no sistema taiga
// @host localhost:8080
// @BasePath /

// @schemes http

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
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

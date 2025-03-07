package main

import (
	"inine-track/pkg/database"
	"inine-track/pkg/routes"
	"log"
)

func main() {

	err := database.ConnectDB()

	if err != nil {
		log.Fatal("erro ao conectar com o banco!", err.Error())
	}

	routes.HandlleRequest()

}

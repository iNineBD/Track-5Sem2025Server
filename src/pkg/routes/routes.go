package routes

import (
	// Import necessário para gerar documentação com Swagger
	_ "inine-track/docs"
	"inine-track/pkg/config"
	"inine-track/pkg/controller"

	"log"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title API Inine-Track
// @version 1.0
// @description Esta é uma API feita para análise de dos projetos no sistema taiga
// @host localhost:8080
// @BasePath /
func HandlleRequest() {
	r := gin.Default()

	r.Use(config.CorsConfig())

	projects := r.Group("/projects")
	{
		projects.GET("/data", controller.GetProjects)
	}

	statistics := r.Group("/statistics")
	{
		statistics.GET("/data/:id", controller.GetStatisticsData)
	}

	// Endpoint Swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	if err := r.Run(); err != nil {
		log.Fatalf("Erro ao iniciar o servidor: %v", err)
	}
}

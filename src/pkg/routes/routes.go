package routes

import (
	"log"

	_ "inine-track/docs" // Necessário para gerar a documentação da API
	"inine-track/pkg/config"
	"inine-track/pkg/controller"
	"inine-track/pkg/middleware"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title API Inine-Track
// @version 1.0
// @description Esta é uma API feita para análise de dos projetos no sistema taiga
// @host localhost:8080
// @BasePath /

// @securityDefinitions.apikey BearerAuth
// @in header
// @type apiKey
// @name Authorization
func HandlleRequest() {
	r := gin.Default()
	r.Use(config.CorsConfig())
	user := r.Group("/access")
	{
		user.POST("/login", controller.LoginController)
		user.POST("/firstAccess", controller.FirstAccessController)
		user.POST("/setPassword", controller.SetPasswordController)
	}

	protected := r.Group("/api")
	protected.Use()
	{

		plataforms := protected.Group("/platforms")
		{
			plataforms.GET("/data", middleware.Auth(), controller.GetPlatforms)
		}
		projects := protected.Group("/projects")
		{
			projects.GET("/data/:idPlatform", middleware.Auth(), controller.GetProjects)
		}
		statistics := protected.Group("/statistics")
		{
			statistics.GET("/data/:id", middleware.Auth(), controller.GetStatisticsData)
		}
		usermanagement := protected.Group("/usermanagement")
		{
			usermanagement.GET("/data", middleware.Auth(), controller.GetRelationUserRole)
			usermanagement.GET("/data/roles", middleware.Auth(), controller.GetRoles)
			usermanagement.PUT("/update", middleware.Auth(), controller.UpdateRoleUser)
		}
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	if err := r.Run(); err != nil {
		log.Fatalf("Erro ao iniciar o servidor: %v", err)
	}
}

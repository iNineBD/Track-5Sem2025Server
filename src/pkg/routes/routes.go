package routes

import (
	"log"

	_ "inine-track/docs"
	"inine-track/pkg/config"
	"inine-track/pkg/controller"
	"inine-track/pkg/middleware"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

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
	protected.Use(middleware.JWTMiddleware())
	{
		projects := r.Group("/projects")
		{
			projects.GET("/data", controller.GetProjects)
		}

		statistics := r.Group("/statistics")
		{
			statistics.GET("/data/:id", controller.GetStatisticsData)
		}
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	if err := r.Run(); err != nil {
		log.Fatalf("Erro ao iniciar o servidor: %v", err)
	}
}

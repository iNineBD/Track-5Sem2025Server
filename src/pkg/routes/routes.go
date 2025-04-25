package routes

import (
	_ "inine-track/docs"
	"inine-track/pkg/config"
	"inine-track/pkg/controller"
	"inine-track/pkg/middleware"
	"inine-track/pkg/service"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func HandlleRequest() {
	r := gin.Default()
	r.Use(config.CorsConfig())

	// Endpoint de login inicial
	r.POST("/login", func(c *gin.Context) {
		var loginRequest struct {
			Email string `json:"email" binding:"required"`
		}

		if err := c.ShouldBindJSON(&loginRequest); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Email é obrigatório"})
			return
		}

		response, err := service.Login(loginRequest.Email)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, response)
	})

	r.POST("/validate-token", func(c *gin.Context) {
		var request struct {
			Email string `json:"email" binding:"required"`
			Token string `json:"token" binding:"required"`
		}

		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Email e token são obrigatórios"})
			return
		}

		valid, err := service.ValidateToken(request.Email, request.Token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		if !valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Token válido"})
	})
	r.POST("/set-password", func(c *gin.Context) {
		var request struct {
			Email       string `json:"email" binding:"required"`
			Token       string `json:"token" binding:"required"`
			NewPassword string `json:"newPassword" binding:"required"`
		}

		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Todos os campos são obrigatórios"})
			return
		}

		if err := service.SetPassword(request.Email, request.Token, request.NewPassword); err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Senha definida com sucesso"})
	})

	r.POST("/authenticate", func(c *gin.Context) {
		var request struct {
			Email    string `json:"email" binding:"required"`
			Password string `json:"password" binding:"required"`
		}

		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Email e senha são obrigatórios"})
			return
		}

		response, err := service.Authenticate(request.Email, request.Password)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, response)
	})

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

package routes

import "github.com/gin-gonic/gin"

// @title API Inine-Track
// @version 1.0
// @description Esta é uma API feita para análise de dos projetos no sistema taiga
// @host localhost:8080
// @BasePath /
func HandlleRequest() {
	r := gin.Default()

	r.Run()
}

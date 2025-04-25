package controller

import (
	"inine-track/pkg/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// LoginHandler realiza o login de um usuário
// @Summary Realiza o login de um usuário
// @Description Autentica o usuário com email e senha e retorna um token JWT
// @Tags Usuário
// @Param email path string true "email do usuario para busca"
// @Produce json
// @Router /login [post]
func LoginHandler(c *gin.Context) {
	var request struct {
		Email string `json:"email"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos: email e senha são obrigatórios"})
		return
	}

	response, err := service.Login(request.Email)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}

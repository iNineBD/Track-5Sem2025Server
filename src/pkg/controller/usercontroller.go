package controller

import (
	"inine-track/pkg/dto/userdto"
	"inine-track/pkg/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Realiza o login de um usuário
// @Description Autentica o usuário com email e senha e retorna um token JWT
// @Tags Usuário
// @Accept json
// @Produce json
// @Param request body userdto.LoginRequest true "Credenciais de login"
// @Router /access/login [post]
func LoginController(c *gin.Context) {
	var loginRequest userdto.LoginRequest
	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Formato de requisição inválido"})
		return
	}

	response, status := service.Login(loginRequest)

	switch status {
	case http.StatusBadRequest:
		c.JSON(status, gin.H{"error": "Usuário não encontrado ou dados inválidos"})
	case http.StatusUnauthorized:
		c.JSON(status, gin.H{"error": "Senha incorreta"})
	case http.StatusInternalServerError:
		c.JSON(status, gin.H{"error": "Falha ao gerar token"})
	default:
		c.JSON(status, gin.H{"success": response})
	}
}

// @Summary Primeiro acesso do usuário
// @Description Gera token para primeiro acesso e envia por email
// @Tags Usuário
// @Accept json
// @Produce json
// @Param request body userdto.FirstAccessRequest true "Email do usuário"
// @Router /access/firstAccess [post]
func FirstAccessController(c *gin.Context) {
	var firstAccessRequest userdto.FirstAccessRequest

	if err := c.ShouldBindJSON(&firstAccessRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Formato de requisição inválido"})
		return
	}

	response, status := service.FirstAccess(firstAccessRequest)

	switch status {
	case http.StatusBadRequest:
		c.JSON(status, gin.H{"error": "Usuário não encontrado ou dados inválidos"})
	case http.StatusInternalServerError:
		c.JSON(status, gin.H{"error": "Falha ao enviar token por email"})
	default:
		c.JSON(status, gin.H{"success": response})
	}
}

// @Summary Define nova senha
// @Description Valida o token e define uma nova senha para o usuário
// @Tags Usuário
// @Accept json
// @Produce json
// @Param request body userdto.AuthenticateRequest true "Token e nova senha"
// @Router /access/setPassword [post]
func SetPasswordController(c *gin.Context) {
	var authenticateRequest userdto.AuthenticateRequest

	if err := c.ShouldBindJSON(&authenticateRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Formato de requisição inválido"})
		return
	}

	status := service.Authenticate(authenticateRequest)

	switch status {
	case http.StatusBadRequest:
		c.JSON(status, gin.H{"error": "Token inválido ou falha ao atualizar senha"})
	case http.StatusUnauthorized:
		c.JSON(status, gin.H{"error": "Token inválido ou expirado"})
	case http.StatusInternalServerError:
		c.JSON(status, gin.H{"error": "Erro interno ao processar requisição"})
	default:
		c.JSON(status, gin.H{"success": "Senha atualizada com sucesso"})
	}
}

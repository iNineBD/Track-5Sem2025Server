package service

import (
	"errors"
	"inine-track/pkg/database"
	"inine-track/pkg/dto/userdto"
	"inine-track/pkg/middleware"
	"inine-track/pkg/models"
	"sync"

	"golang.org/x/crypto/bcrypt"
)

type TokenStore struct {
	mu    sync.Mutex
	store map[string]string // email -> token
}

var tokenStore = TokenStore{
	store: make(map[string]string),
}

func Login(email, password string) (userdto.GetUserResponse, error) {
	var user models.DimUser
	response := userdto.GetUserResponse{}

	// Busca o usuário
	result := database.DB.
		Preload("DimRole").
		Where("email = ?", email).
		First(&user)

	if result.Error != nil {
		return response, errors.New("usuário não encontrado")
	}

	// Se o usuário não tem senha definida
	if user.Password == "" {
		token, err := middleware.GenerateJWT(user.IDUser, user.Email, user.IDRole)
		if err != nil {
			return response, errors.New("erro ao gerar token")
		}

		// Armazena o token
		tokenStore.mu.Lock()
		tokenStore.store[email] = token
		tokenStore.mu.Unlock()

		// Envia o email com o token
		if err := SendTokenEmail(email, token); err != nil {
			return response, errors.New("erro ao enviar email com token")
		}
		return response, nil
	}

	// Verifica a senha para usuários com senha definida
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return response, errors.New("credenciais inválidas")
	}

	// Gera token para login bem-sucedido
	token, err := middleware.GenerateJWT(user.IDUser, user.Email, user.IDRole)
	if err != nil {
		return response, errors.New("erro ao gerar token de autenticação")
	}

	response.ID = user.IDUser
	response.Email = user.Email
	response.Role = user.IDRole
	response.Token = token

	return response, nil
}

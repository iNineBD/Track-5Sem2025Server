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

func Login(email string) (userdto.GetUserResponse, error) {
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

	if user.Password == "" {
		token, err := middleware.GenerateJWT(user.IDUser, user.Email, user.IDRole)
		if err != nil {
			return response, errors.New("erro ao gerar token")
		}

		tokenStore.mu.Lock()
		tokenStore.store[email] = token
		tokenStore.mu.Unlock()

		if err := SendTokenEmail(email, token); err != nil {
			return response, errors.New("erro ao enviar email com token")
		}

		return response, nil
	}

	response.ID = user.IDUser
	response.Email = user.Email
	response.Role = user.IDRole

	return response, nil
}

func ValidateToken(email, token string) (bool, error) {
	tokenStore.mu.Lock()
	defer tokenStore.mu.Unlock()

	storedToken, exists := tokenStore.store[email]
	if !exists {
		return false, errors.New("token não encontrado ou expirado")
	}

	return token == storedToken, nil
}

func SetPassword(email, token, newPassword string) error {
	valid, err := ValidateToken(email, token)
	if err != nil {
		return err
	}
	if !valid {
		return errors.New("token inválido")
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("erro ao encriptar senha")
	}
	result := database.DB.Model(&models.DimUser{}).
		Where("email = ?", email).
		Update("password", string(hashedPassword))

	if result.Error != nil {
		return errors.New("erro ao atualizar senha no banco de dados")
	}

	tokenStore.mu.Lock()
	delete(tokenStore.store, email)
	tokenStore.mu.Unlock()

	return nil
}

func Authenticate(email, password string) (userdto.GetUserResponse, error) {
	var user models.DimUser
	response := userdto.GetUserResponse{}

	result := database.DB.
		Preload("DimRole").
		Where("email = ?", email).
		First(&user)

	if result.Error != nil {
		return response, errors.New("usuário não encontrado")
	}

	if user.Password == "" {
		return response, errors.New("primeiro acesso requerido - verifique seu email")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return response, errors.New("credenciais inválidas")
	}

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

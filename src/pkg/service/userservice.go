package service

import (
	"errors"
	"inine-track/pkg/database"
	"inine-track/pkg/dto/userdto"
	"inine-track/pkg/middleware"
	"inine-track/pkg/models"
	"net/http"
	"sync"

	"golang.org/x/crypto/bcrypt"
)

func Login(request userdto.LoginRequest) (userdto.GetUserResponse, int) {
	var user models.DimUser
	response := userdto.GetUserResponse{}

	result := database.DB.
		Preload("DimRole").
		Where("email = ?", request.Email).
		First(&user)

	if result.Error != nil {
		return response, http.StatusBadRequest
	}

	if user.Password == "" {
		return response, http.StatusBadRequest
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)); err != nil {
		return response, http.StatusBadRequest
	}

	token, err := middleware.GenerateJWT(user.IDUser, user.Email, user.IDRole)
	if err != nil {
		return response, http.StatusBadRequest
	}

	response.ID = user.IDUser
	response.Email = user.Email
	response.Role = user.IDRole
	response.Token = token

	return response, http.StatusOK
}

type TokenStore struct {
	mu    sync.Mutex
	store map[string]string
}

var tokenStore = TokenStore{
	store: make(map[string]string),
}

func FirstAccess(request userdto.FirstAccessRequest) (userdto.GetUserResponse, int) {
	var user models.DimUser
	response := userdto.GetUserResponse{}

	// Busca o usuário
	result := database.DB.
		Preload("DimRole").
		Where("email = ?", request.Email).
		First(&user)

	if result.Error != nil {
		return response, http.StatusBadRequest
	}

	if user.Password == "" {
		token, err := middleware.GenerateJWT(user.IDUser, user.Email, user.IDRole)
		if err != nil {
			return response, http.StatusBadRequest
		}

		tokenStore.mu.Lock()
		tokenStore.store[request.Email] = token
		tokenStore.mu.Unlock()

		if err := SendTokenEmail(request.Email, token); err != nil {
			return response, http.StatusBadRequest
		}

		return response, http.StatusOK
	}

	response.ID = user.IDUser
	response.Email = user.Email
	response.Role = user.IDRole

	return response, http.StatusOK
}

func ValidateToken(request userdto.AuthenticateRequest) (bool, error) {
	tokenStore.mu.Lock()
	defer tokenStore.mu.Unlock()

	storedToken, exists := tokenStore.store[request.Email]
	if !exists {
		return false, errors.New("token não encontrado ou expirado")
	}

	return request.Token == storedToken, nil
}

func Authenticate(request userdto.AuthenticateRequest) int {
	valid, err := ValidateToken(request)
	if err != nil {
		return http.StatusBadRequest
	}
	if !valid {
		return http.StatusBadRequest
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		return http.StatusBadRequest
	}
	result := database.DB.Model(&models.DimUser{}).
		Where("email = ?", request.Email).
		Update("password", string(hashedPassword))

	if result.Error != nil {
		return http.StatusBadRequest
	}

	tokenStore.mu.Lock()
	delete(tokenStore.store, request.Email)
	tokenStore.mu.Unlock()

	return http.StatusOK
}

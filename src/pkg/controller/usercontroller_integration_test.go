package controller

// Importação dos pacotes necessários
import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"inine-track/pkg/database"
	"inine-track/pkg/dto/userdto"
	"inine-track/pkg/models"
)

// Função auxiliar que gera um token JWT válido para testes.
// Isso simula a geração de token como feito em produção.
func generateTestJWT(userID int64, email string, role int64) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"email":   email,
		"role":    role,
		"exp":     time.Now().Add(1 * time.Hour).Unix(), // Expiração do token
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte("test-secret-key-for-integration-tests"))
}

// Controller de login simulado para testes de integração.
// Interage com o banco real (sqlite em memória) e os modelos reais DimUser e DimRole.
func MockLoginController(c *gin.Context) {
	var loginRequest userdto.LoginRequest
	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Formato de requisição inválido"})
		return
	}

	var user models.DimUser
	result := database.DB.
		Preload("DimRole"). // Traz junto os dados do relacionamento com a tabela DimRole
		Where("email = ?", loginRequest.Email).
		First(&user)

	// Logs úteis para depuração
	fmt.Printf("Login attempt for email: %s\n", loginRequest.Email)
	if result.Error != nil {
		fmt.Printf("DB Error: %v\n", result.Error)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Usuário não encontrado ou dados inválidos"})
		return
	}

	fmt.Printf("User found: ID=%d, Email=%s\n", user.IDUser, user.Email)

	if user.Password == "" {
		fmt.Println("Empty password")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Usuário não encontrado ou dados inválidos"})
		return
	}

	// Validação da senha com bcrypt
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginRequest.Password)); err != nil {
		fmt.Printf("Password comparison error: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Usuário não encontrado ou dados inválidos"})
		return
	}

	// Geração do token JWT
	token, err := generateTestJWT(user.IDUser, user.Email, user.IDRole)
	if err != nil {
		fmt.Printf("JWT generation error: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Falha ao gerar token"})
		return
	}

	// Resposta com os dados do usuário e o token JWT
	response := userdto.GetUserResponse{
		ID:       user.IDUser,
		NameUser: strings.ToUpper(user.NameUser),
		Email:    user.Email,
		Role:     user.IDRole,
		NameRole: strings.ToUpper(user.DimRole.NameRole),
		Token:    token,
	}

	c.JSON(http.StatusOK, gin.H{"success": response})
}

// Define variáveis de ambiente necessárias para os testes
func setupTestEnvironment() {
	os.Setenv("JWT_SECRET", "test-secret-key-for-integration-tests")
}

// Teste de integração completo do controller de login
func TestLoginControllerIntegration(t *testing.T) {
	gin.SetMode(gin.TestMode)

	setupTestEnvironment()

	// Cria um banco SQLite temporário e reseta os dados antes de cada teste
	resetDB := func() *gorm.DB {
		dir := os.TempDir()
		dbFile, err := os.CreateTemp(dir, "testdb_*.db")
		assert.NoError(t, err, "Falha ao criar arquivo temporário para o banco de dados")
		db, err := gorm.Open(sqlite.Open(dbFile.Name()), &gorm.Config{})
		assert.NoError(t, err, "Falha ao conectar ao banco de dados temporário")
		err = db.AutoMigrate(&models.DimUser{}, &models.DimRole{}) // Migração das tabelas
		assert.NoError(t, err, "Falha ao migrar o schema do banco de dados")
		database.DB = db
		return db
	}

	// Teste: login válido com dados reais
	t.Run("Valid Login - Integration", func(t *testing.T) {
		db := resetDB()
		// Cria role e usuário reais
		role := models.DimRole{IDRole: 1, NameRole: "ADMIN"}
		result := db.Create(&role)
		assert.NoError(t, result.Error, "Failed to create role")

		password := "securepassword123"
		hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		assert.NoError(t, err, "Failed to hash password")

		testUser := models.DimUser{
			NameUser: "Integration User",
			Email:    "integration@example.com",
			Password: string(hash),
			IDRole:   role.IDRole,
		}
		result = db.Create(&testUser)
		assert.NoError(t, result.Error, "Failed to create user")

		// Verifica se o preload da role funciona
		var checkUser models.DimUser
		result = db.Preload("DimRole").Where("email = ?", testUser.Email).First(&checkUser)
		assert.NoError(t, result.Error, "Failed to retrieve created user")

		// Simula uma requisição HTTP POST para o endpoint /access/login
		router := gin.Default()
		router.POST("/access/login", MockLoginController)

		loginRequest := userdto.LoginRequest{
			Email:    testUser.Email,
			Password: password,
		}
		body, _ := json.Marshal(loginRequest)
		req, _ := http.NewRequest("POST", "/access/login", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)

		// Valida a resposta do controller
		assert.Equal(t, http.StatusOK, resp.Code)

		var response map[string]interface{}
		err = json.Unmarshal(resp.Body.Bytes(), &response)
		assert.NoError(t, err)

		success, ok := response["success"].(map[string]interface{})
		assert.True(t, ok, "Response does not contain a success object")

		if ok {
			token, tokenExists := success["token"]
			assert.True(t, tokenExists, "Token field not found in success object")
			assert.NotEmpty(t, token, "Token should not be empty")
		}
	})

	// Teste: requisição mal formatada
	t.Run("Invalid Login Format - Integration", func(t *testing.T) {
		resetDB()
		router := gin.Default()
		router.POST("/access/login", MockLoginController)
		invalidBody := []byte(`{"invalid": "data"}`)
		req, _ := http.NewRequest("POST", "/access/login", bytes.NewBuffer(invalidBody))
		req.Header.Set("Content-Type", "application/json")
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusBadRequest, resp.Code)
		assert.Contains(t, resp.Body.String(), "Formato de requisição inválido")
	})

	// Teste: usuário não encontrado
	t.Run("Unauthorized Login - User Not Found - Integration", func(t *testing.T) {
		resetDB()
		router := gin.Default()
		router.POST("/access/login", MockLoginController)
		loginRequest := userdto.LoginRequest{
			Email:    "nonexistent@example.com",
			Password: "wrongpassword",
		}
		body, _ := json.Marshal(loginRequest)
		req, _ := http.NewRequest("POST", "/access/login", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusBadRequest, resp.Code)
		assert.Contains(t, resp.Body.String(), "Usuário não encontrado ou dados inválidos")
	})

	// Teste: senha incorreta
	t.Run("Unauthorized Login - Incorrect Password - Integration", func(t *testing.T) {
		db := resetDB()
		role := models.DimRole{IDRole: 1, NameRole: "ADMIN"}
		db.Create(&role)
		password := "correctpassword"
		hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		testUser := models.DimUser{
			NameUser: "Password User",
			Email:    "passwordcheck@example.com",
			Password: string(hash),
			IDRole:   role.IDRole,
		}
		db.Create(&testUser)

		router := gin.Default()
		router.POST("/access/login", MockLoginController)

		loginRequest := userdto.LoginRequest{
			Email:    testUser.Email,
			Password: "incorrectpassword",
		}
		body, _ := json.Marshal(loginRequest)
		req, _ := http.NewRequest("POST", "/access/login", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusBadRequest, resp.Code)
		assert.Contains(t, resp.Body.String(), "Usuário não encontrado ou dados inválidos")
	})
}

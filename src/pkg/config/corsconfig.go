package config

import (
	"errors"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type EmailConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	From     string
}

var SMTP *EmailConfig

func CorsConfig() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "http://144.22.212.19:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           3 * time.Hour,
	})
}

func LoadJWTKey() (string, error) {
	err := godotenv.Load()
	if err != nil {
		return "", err
	}
	return os.Getenv("JWT_SECRET"), nil
}

func LoadSendEmail() error {
	err := godotenv.Load()
	if err != nil {
		return err
	}
	SMTP = &EmailConfig{
		Host:     os.Getenv("EMAIL_HOST"),
		Port:     587,
		User:     os.Getenv("EMAIL_HOST_USERNAME"),
		Password: os.Getenv("EMAIL_HOST_PASSWORD"),
		From:     os.Getenv("EMAIL_HOST_FROM"),
	}
	if SMTP.Host == "" || SMTP.User == "" || SMTP.Password == "" {
		return errors.New("SMTP configuration is incomplete")
	}
	return nil

}

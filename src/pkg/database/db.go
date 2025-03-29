package database

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() (err error) {

	paths := []string{".env", "../../.env", "src/.env", "../.env"}
	var loadErr error

	for _, path := range paths {
		if err := godotenv.Load(path); err == nil {
			loadErr = nil
			break
		} else {
			loadErr = err
		}
	}

	if loadErr != nil {
		return fmt.Errorf("erro ao carregar o arquivo .env : %s", loadErr.Error())
	}

	// Obter as variáveis de ambiente
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	name := os.Getenv("DB_NAME")
	schema := os.Getenv("DB_SCHEMA")

	// Montar DSN (Data Source Name)
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable search_path=%s", host, user, password, name, port, schema)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("erro ao conectar ao banco: %s", err.Error())
	}
	fmt.Println("Conexão bem-sucedida ao banco Postgres!")

	return nil
}

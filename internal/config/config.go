package config

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"

	"github.com/joho/godotenv"
)

var JwtKey []byte = []byte("BandidoSecretPassword123?_45")

type Config struct {
	DB  PostgreConfig
	JWT JwtConfig
}
type PostgreConfig struct {
	Username string
	Password string
	Host     string
	DName    string
	Port     int
}

type JwtConfig struct {
	JwtSecret string
}

func LoadConfig() (*Config, error) {
	baseDir, err := filepath.Abs("../..")
	if err != nil {
		return nil, err
	} // Subir 3 directorios desde cmd/app
	envPath := filepath.Join(baseDir, "internal", "config")
	err = godotenv.Load(envPath, ".env")
	if err != nil {
		return nil, errors.New(".env file not found or could not be loaded")
	}

	portEnv := os.Getenv("DBPORT")
	User := os.Getenv("DBUSERNAME")
	fmt.Printf("DB user: %s\n", User)
	if portEnv == "" {
		portEnv = "5433"
		log.Printf("PORT environment variable is not set, using default port: %s", portEnv)
	}

	value, errorInt := strconv.Atoi(portEnv)
	if errorInt != nil {
		return nil, err
	}
	cfg := &Config{
		DB: PostgreConfig{
			Username: os.Getenv("DBUSERNAME"),
			Password: os.Getenv("DBPASSWORD"),
			Host:     os.Getenv("DBHOST"),
			DName:    os.Getenv("DBDATASOURCE"),
			Port:     value,
		},
		JWT: JwtConfig{
			JwtSecret: os.Getenv("SECRET_TOKEN_KEY"),
		},
	}
	return cfg, nil
}

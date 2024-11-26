package service

import (
	"github/1AdrianM/go_inventario/internal/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func CheckPassword(HashedPassword, password string) bool {
	// HashedPassword is the hashed password from the database
	// password is the password the user entered
	err := bcrypt.CompareHashAndPassword([]byte(HashedPassword), []byte(password))
	return nil == err
}

func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func GenerateJWT(email string, userID uint) (string, error) {

	claims := &jwt.MapClaims{
		"email":  email,
		"UserID": userID,
		"exp":    time.Now().Add(24 * time.Hour).Unix(), // Token expira en 24 horas
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString(config.JwtKey)
	if err != nil {
		return "", err
	}
	return tokenStr, nil
}

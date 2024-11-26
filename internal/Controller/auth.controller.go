package controller

import (
	"errors"
	model "github/1AdrianM/go_inventario/internal/Model"
	repository "github/1AdrianM/go_inventario/internal/Repository"
	service "github/1AdrianM/go_inventario/internal/Service"
	"log"
)

func LogIn(email string, password string) (user *model.User, err error) {

	user, err = repository.GetUserByEmail(email)

	if err != nil || user == nil || !service.CheckPassword(user.Password, password) {
		return nil, errors.New("invalid credentials")
	}
	return user, err
}

func SignUp(user *model.User) error {
	// hashear contrasena de usuario

	hashedPasword, err := service.HashPassword(user.Password)
	if err != nil {
		return errors.New("not able to hash password")
	}
	//crear objecto user tipo User
	usr := model.User{Name: user.Name, Lastname: user.Lastname, Email: user.Email, Password: hashedPasword}
	// Insertar el user en la base de datos
	if err := repository.CreateUser(&usr); err != nil {

		// Para otros errores, redirigir con un mensaje genérico
		log.Printf("Error creating user: %v", err)
		return errors.New("error creating user")
	}
	return nil
}

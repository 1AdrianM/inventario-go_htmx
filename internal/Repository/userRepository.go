package repository

import (
	"errors"
	db "github/1AdrianM/go_inventario/internal/Db"
	model "github/1AdrianM/go_inventario/internal/Model"
)

func GetUserById(id uint) (*model.User, error) {
	var user model.User

	if err := db.Db.First(&user, id).Error; err != nil {
		return nil, errors.New("user not found")

	}
	return &user, nil
}
func GetUserByEmail(email string) (*model.User, error) {
	var user model.User

	if err := db.Db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, errors.New("user not found")

	}
	return &user, nil
}
func GetUsers() ([]model.User, error) {
	var users []model.User

	if err := db.Db.Find(&users).Error; err != nil {
		return nil, errors.New("failed to retrieve users")
	}
	return users, nil

}
func CreateUser(user *model.User) error {
	if user.Name == "" || user.Email == "" || user.Password == "" || user.Lastname == "" {
		return errors.New("input data is required")
	}
	var existingUser model.User

	result := db.Db.Where(" Email=?", user.Email).First(&existingUser)

	if result.RowsAffected > 0 {
		return errors.New("user already exists")
	}

	if err := db.Db.Create(&user).Error; err != nil {
		return errors.New("failed to create user")
	}

	return nil
}

func DeleteUser(id int) error {
	if err := db.Db.Delete(&model.User{}, id).Error; err != nil {
		return errors.New("failed to delete user")

	}
	return nil
}

func UpdateUser(id int, UpdateUser *model.User) error {
	var user model.User

	if err := db.Db.First(&user, id).Error; err != nil {
		return errors.New("user not found")

	}
	if UpdateUser.Email != "" {
		user.Email = UpdateUser.Email
	}
	if UpdateUser.Name != "" {
		user.Name = UpdateUser.Name
	}
	if UpdateUser.Lastname != "" {
		user.Lastname = UpdateUser.Lastname
	}
	if UpdateUser.Password != "" {
		user.Password = UpdateUser.Password
	}
	if err := db.Db.Save(&user).Error; err != nil {
		return err
	}
	return nil
}

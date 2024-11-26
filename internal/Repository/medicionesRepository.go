package repository

import (
	"errors"
	db "github/1AdrianM/go_inventario/internal/Db"
	model "github/1AdrianM/go_inventario/internal/Model"
	"time"
)

func CreateMediciones(medicion *model.Mediciones) error {
	if medicion.Fecha_mantenimiento == (time.Time{}) ||
		medicion.Motor_ubicacion == "" ||
		medicion.Motor_KW == 0 ||
		medicion.Motor_RPM == 0 ||
		medicion.Motor_tag == "" ||
		medicion.Bearing_d == 0 ||
		medicion.Bearing_t == 0 ||
		medicion.UserID == 0 {
		return errors.New("invalid data")
	}
	var user model.User

	fetchUser := db.Db.First(&user, medicion.UserID)
	if fetchUser.RowsAffected == 0 {
		return errors.New("not found a user with ID")

	}
	result := db.Db.Create(&medicion)
	if result.Error != nil {
		return errors.New("not been able to create the medicion")
	}

	return nil
}

func GetMediciones() ([]model.Mediciones, error) {
	var mediciones []model.Mediciones

	if err := db.Db.Find(&mediciones).Error; err != nil {

		return nil, errors.New("couldnt find mediciones")

	}
	return mediciones, nil
}

func GetMedicionById(id uint) (*model.Mediciones, error) {

	var medicion model.Mediciones

	if err := db.Db.First(&medicion, id).Error; err != nil {
		return nil, errors.New("couldnt find medicion")
	}
	return &medicion, nil
}

func DeleteMediciones(id uint) error {
	if id == 0 {
		return errors.New("ID not found")
	}
	if err := db.Db.Delete(&model.Mediciones{}, id).Error; err != nil {
		return errors.New("couldnt delete medicion")
	}
	return nil
}

func UpdateMediciones(id uint, UpdateMediciones *model.Mediciones) error {
	var medicion model.Mediciones
	if err := db.Db.First(&medicion, id).Error; err != nil {
		return errors.New("couldnt find medicion")
	}
	if UpdateMediciones.Fecha_mantenimiento != (time.Time{}) {
		medicion.Fecha_mantenimiento = UpdateMediciones.Fecha_mantenimiento
	}
	if UpdateMediciones.Motor_ubicacion != "" {
		medicion.Motor_ubicacion = UpdateMediciones.Motor_ubicacion
	}
	if UpdateMediciones.Motor_KW != 0 {
		medicion.Motor_KW = UpdateMediciones.Motor_KW
	}
	if UpdateMediciones.Motor_RPM != 0 {
		medicion.Motor_RPM = UpdateMediciones.Motor_RPM
	}
	if UpdateMediciones.Motor_tag != "" {
		medicion.Motor_tag = UpdateMediciones.Motor_tag
	}
	if UpdateMediciones.Bearing_d != 0 {
		medicion.Bearing_d = UpdateMediciones.Bearing_d
	}
	if UpdateMediciones.Bearing_t != 0 {
		medicion.Bearing_t = UpdateMediciones.Bearing_t
	}

	if err := db.Db.Save(&medicion).Error; err != nil {
		return errors.New("couldnt update medicion")
	}
	return nil
}

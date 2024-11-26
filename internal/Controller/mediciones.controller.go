package controller

import (
	"errors"
	model "github/1AdrianM/go_inventario/internal/Model"
	repository "github/1AdrianM/go_inventario/internal/Repository"
	"github/1AdrianM/go_inventario/internal/helper"
)

func GetAllMediciones() ([]model.Mediciones, error) {
	helper.Info("GetAllMediciones function called")

	result, err := repository.GetMediciones()
	if err != nil {
		return nil, errors.New("not found mediciones")
	}
	return result, nil
}

func GetMedicionById(id uint) (*model.Mediciones, error) {
	mediciones, err := repository.GetMedicionById(id)
	if err != nil {
		return nil, errors.New("not found medicion")
	}
	return mediciones, nil
}

func CreateMedicion(medicion *model.Mediciones, user *model.User) error {
	if err := repository.CreateMediciones(medicion); err != nil {
		return errors.New("could not create medicion")
	}
	return nil
}
func UpdateMedicion(medicion model.Mediciones) {

}
func DeleteMedicion(id uint) error {

	if err := repository.DeleteMediciones(id); err != nil {
		return errors.New("could not delete medicion")
	}
	return nil
}

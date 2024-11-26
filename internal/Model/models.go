package model

import (
	"time"

	"gorm.io/gorm"
)

type Mediciones struct {
	gorm.Model
	Fecha_mantenimiento time.Time
	Motor_ubicacion     string
	Motor_tag           string
	Motor_KW            int
	Motor_RPM           int
	Bearing_d           float64
	Bearing_t           float64
	Bearing_a           float64

	UserID uint
	User   User `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE;"` // Explicit foreign key definition

}
type User struct {
	gorm.Model
	Name       string
	Lastname   string
	Email      string
	Password   string
	Role       string
	Mediciones []Mediciones
	Task       []Task
}

type Task struct {
	gorm.Model
	Title     string
	Text      string
	Completed bool   `gorm:"default:false"`
	Status    string `gorm:"default:'In Progress'"`
	Due_Date  time.Time
	UserID    uint // Clave foránea para User
	User      User `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE;"`
}

package domain

import "gorm.io/gorm"

type Admin struct {
	gorm.Model
	Name     string
	Password string
	Email    string
}

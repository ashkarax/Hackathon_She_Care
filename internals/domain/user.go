package domain

import "gorm.io/gorm"

type Users struct {
	gorm.Model
	Name     string
	Password string
	Email    string
}

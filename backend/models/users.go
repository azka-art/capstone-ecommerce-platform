package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string    `json:"username"`
	Password string    `json:"password"`
	Products []Product `gorm:"many2many:transactions;"`
}

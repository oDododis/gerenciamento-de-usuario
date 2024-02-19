package model

import "gorm.io/gorm"

// Estrutura do usuário

type User struct {
	gorm.Model
	FullName string
	Email    string
	Username string
	Password string
}

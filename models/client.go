package models

import "gorm.io/gorm"

type Client struct {
	gorm.Model
	Name  string
	Email string
	Cpf   string
}

package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	ClientID uint
	Client   Client
	Status   int8
	Total    float64
}

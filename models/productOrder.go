package models

import "gorm.io/gorm"

type ProductOrder struct {
	gorm.Model
	OrderID   uint
	ProductID uint
}

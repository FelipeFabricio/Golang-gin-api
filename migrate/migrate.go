package main

import (
	"github.com/felipefabricio/golang-gin-api/initializers"
	"github.com/felipefabricio/golang-gin-api/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	_ = initializers.DB.AutoMigrate(&models.Client{}, &models.Category{}, &models.Product{}, &models.Order{}, &models.ProductOrder{})
}

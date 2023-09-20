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
	initializers.DB.AutoMigrate(&models.Client{})
}

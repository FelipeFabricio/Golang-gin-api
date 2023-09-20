package controllers

import (
	"github.com/felipefabricio/golang-gin-api/initializers"
	"github.com/felipefabricio/golang-gin-api/models"
	"github.com/gin-gonic/gin"
)

func ClientCreate(c *gin.Context) {
	var body struct {
		Name  string `json:"name"`
		Email string `json:"email"`
		Cpf   string `json:"cpf"`
	}

	c.Bind(&body)
	client := models.Client{Name: body.Name, Email: body.Email, Cpf: body.Cpf}
	result := initializers.DB.Create(&client)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"client": body,
	})
}

func ClientList(c *gin.Context) {
	var clients []models.Client
	initializers.DB.Find(&clients)

	c.JSON(200, gin.H{
		"clients": clients,
	})
}

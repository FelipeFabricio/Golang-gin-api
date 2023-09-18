package main

import (
	"fmt"
	"os"

	"github.com/felipefabricio/golang-gin-api/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()
	fmt.Printf("Server running on port %s", os.Getenv("PORT"))
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Hello World"})
	})
	r.Run()
}

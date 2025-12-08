package main


	import (
    "github.com/mareligncode/selam/config"
    "github.com/mareligncode/selam/models"
    "github.com/mareligncode/selam/routes"
     "github.com/gin-gonic/gin"
)


func main() {
	// Connect MySQL
	config.Connect()

	// AutoMigrate Models
	config.AutoMigrate(&models.User{})

	r := gin.Default()

	// Load Bootstrap HTML templates
	r.LoadHTMLGlob("templates/*")

	// Load CSS
	r.Static("/static", "./static")

	// Register routes
	routes.RegisterRoutes(r)

	r.Run(":8080")
}

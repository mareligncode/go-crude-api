package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mareligncode/selam/config"
	"github.com/mareligncode/selam/models"
)

// Get all users - returns JSON
func GetUsers(c *gin.Context) {
	var users []models.User
	config.DB.Find(&users)
	c.JSON(http.StatusOK, users)
}

// Get single user by ID
func GetUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User
	if err := config.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, user)
}

// Create new user from JSON
func CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Create(&user)
	c.JSON(http.StatusCreated, user)
}

// Update user by ID from JSON
func UpdateUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User

	if err := config.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	var input models.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user.Name = input.Name
	user.Email = input.Email
	user.Phone = input.Phone

	config.DB.Save(&user)
	c.JSON(http.StatusOK, user)
}

// Delete user by ID
func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User
	if err := config.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	config.DB.Delete(&user)
	c.JSON(http.StatusOK, gin.H{"message": "User deleted"})
}

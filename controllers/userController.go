package controllers

import (
    "github.com/mareligncode/selam/config"
    "github.com/mareligncode/selam/models"
    "github.com/gin-gonic/gin"
    "net/http"
)


func UserIndex(c *gin.Context) {
	var users []models.User
	config.DB.Find(&users)

	c.HTML(http.StatusOK, "index.html", gin.H{
		"users": users,
	})
}

func UserCreatePage(c *gin.Context) {
	c.HTML(http.StatusOK, "create.html", nil)
}

func UserCreate(c *gin.Context) {
	user := models.User{
		Name:  c.PostForm("name"),
		Email: c.PostForm("email"),
		Phone: c.PostForm("phone"),
	}
	config.DB.Create(&user)

	c.Redirect(http.StatusSeeOther, "/")
}

func UserEditPage(c *gin.Context) {
	id := c.Param("id")
	var user models.User

	config.DB.First(&user, id)

	c.HTML(http.StatusOK, "edit.html", gin.H{
		"user": user,
	})
}

func UserUpdate(c *gin.Context) {
	id := c.PostForm("id")

	var user models.User
	config.DB.First(&user, id)

	user.Name = c.PostForm("name")
	user.Email = c.PostForm("email")
	user.Phone = c.PostForm("phone")

	config.DB.Save(&user)

	c.Redirect(http.StatusSeeOther, "/")
}

func UserDelete(c *gin.Context) {
	id := c.Param("id")

	var user models.User
	config.DB.Delete(&user, id)

	c.Redirect(http.StatusSeeOther, "/")
}

package routes

import (
	"github.com/mareligncode/selam/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {

	r.GET("/", controllers.UserIndex)

	r.GET("/create", controllers.UserCreatePage)
	r.POST("/create", controllers.UserCreate)

	r.GET("/edit/:id", controllers.UserEditPage)
	r.POST("/update", controllers.UserUpdate)

	r.GET("/delete/:id", controllers.UserDelete)
}

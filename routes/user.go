package routes

import (
	"go-auth/controller"

	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine) {
	router.GET("/", controller.GetUsers)
	router.POST("/", controller.CreateUser)
	router.POST("/:id", controller.DeleteUser)
	router.PUT("/:id", controller.UpdateUser)
}

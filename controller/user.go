package controller

import (
	"go-auth/config"
	"go-auth/models"

	"github.com/gin-gonic/gin"
)

func UserController(c *gin.Context) {
	users := []models.User{}
	config.DB.Find(&users)
	c.JSON(200, &users)
}

package controller

import (
	"go-auth/config"
	"go-auth/models"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	users := []models.User{}
	config.DB.Find(&users)
	c.JSON(200, &users)
}

func CreateUser(c *gin.Context) {

}

func DeleteUser(c *gin.Context) {

}

func UpdateUser(c *gin.Context) {
	
}
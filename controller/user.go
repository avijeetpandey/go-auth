package controller

import (
	"go-auth/config"
	"go-auth/models"

	"github.com/gin-gonic/gin"
)

// controller to get the list of users
func GetUsers(c *gin.Context) {
	users := []models.User{}
	config.DB.Find(&users)
	c.JSON(200, &users)
}

// controller to create the user
func CreateUser(c *gin.Context) {
	var user models.User
	c.BindJSON(&user)
	config.DB.Create(&user)
	c.JSON(200, &user)
}

// controller to delete the user
func DeleteUser(c *gin.Context) {
	var user models.User
	config.DB.Where("id = ?", c.Param("id")).Delete(&user)
	c.JSON(200, &user)
}

// controller to update the user
func UpdateUser(c *gin.Context) {
	var user models.User
	config.DB.Where("id = ?", c.Param("id")).First((&user))
	c.BindJSON(&user)
	config.DB.Save(&user)
	c.JSON(200, &user)
}

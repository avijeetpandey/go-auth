package main

import (
	"go-auth/config"
	"go-auth/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()
	config.Connect()
	routes.UserRoute(router)
	router.Run(":3000")
}

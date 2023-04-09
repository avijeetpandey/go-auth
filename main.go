package main

import (
	"go-auth/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()

	// routes
	routes.UserRoute(router)

	router.Run(":3000")
}

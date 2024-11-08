package main

import (
	_ "import-swag/docs"
	"import-swag/internal/handler"

	"github.com/gin-gonic/gin"
)

func main() {

	handler := &handler.UserHandler{}

	router := gin.Default()
	api := router.Group("/api")

	api.GET("/users/:userId", handler.GetUser)

	router.Run(":8000")
}

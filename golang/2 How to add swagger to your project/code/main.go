package main

import (
	_ "import-swag/docs"
	"net/http"

	"github.com/gin-gonic/gin"
	files "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// @title import-swag-api
// @version 1.0
// @BasePath /api
func main() {
	router := gin.Default()
	api := router.Group("/api")

	api.GET("/users", userHandler)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(files.Handler))

	router.Run(":8000")
}

// @Summary Get all users
// @Tags users
// @Success 200 {array} User
// @Router /users [get]
func userHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, []User{
		{
			ID:   "4535345843",
			Name: "John",
		},
		{
			ID:   "4535345843",
			Name: "Bob",
		},
	})
}

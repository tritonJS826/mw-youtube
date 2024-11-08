package handler

import (
	"import-swag/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service service.UserService
}

func (uh *UserHandler) GetUser(c *gin.Context) {
	userID := c.Param("userId")
	user, err := uh.service.GetUser(userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
	}
	c.JSON(http.StatusOK, user)
}

package handler

import (
	"facade/internal/facade"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	facade facade.UserFacade
}

func (uh *UserHandler) GetUser(c *gin.Context) {
	userID := c.Param("userId")
	user, err := uh.facade.GetUserDetails(userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
	}
	c.JSON(http.StatusOK, user)
}

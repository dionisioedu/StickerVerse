package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateUserHandler(c *gin.Context) {
	authUser, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	u, ok := authUser.(*User)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user context"})
		return
	}

	var req struct {
		Display string `json:"display" binding:"max=20"`
		Bio     string `json:"bio" binding:"max=1000"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid display name"})
		return
	}

	if err := UpdateUser(u.ID, req.Display, req.Bio); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update display name"})
		return
	}

	u.Display = &req.Display
	u.Bio = &req.Bio
	c.JSON(http.StatusOK, gin.H{"display": u.Display, "bio": u.Bio})
}

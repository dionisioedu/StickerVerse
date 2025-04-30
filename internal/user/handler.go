package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateUserDisplayHandler(c *gin.Context) {
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
		Display string `json:"display" binding:"required,min=2,max=50"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid display name"})
		return
	}

	if err := UpdateUserDisplay(u.ID, req.Display); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update display name"})
		return
	}

	u.Display = req.Display
	c.JSON(http.StatusOK, gin.H{"display": u.Display})
}

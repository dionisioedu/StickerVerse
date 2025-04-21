package auth

import (
	"net/http"

	"github.com/dionisioedu/StickerVerse/internal/user"
	"github.com/gin-gonic/gin"
)

type GoogleAuthRequest struct {
	Token string `json:"token" binding:"required"`
}

func GoogleAuthHandler(c *gin.Context) {
	var req GoogleAuthRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Token is required"})
		return
	}

	googleUser, err := VerifyGoogleToken(req.Token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Google token"})
		return
	}

	u, err := user.FindOrCreateUser(googleUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	token, err := GenerateJWT(u)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
		"user":  u,
	})
}

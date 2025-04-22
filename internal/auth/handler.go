package auth

import (
	"log"
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
		log.Println("Failed to bind Google auth request:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Token is required"})
		return
	}

	log.Println("Verifying Google ID token: " + req.Token)

	googleUser, err := VerifyGoogleToken(req.Token)
	if err != nil {
		log.Println("Google token verification failed:", err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Google token"})
		return
	}

	log.Printf("Google token verified. User: %+v\n", googleUser)

	u, err := user.FindOrCreateUser(googleUser)
	if err != nil {
		log.Println("Failed to find or create user:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	log.Printf("User authenticated: %+v\n", u)

	token, err := GenerateJWT(u)
	if err != nil {
		log.Println("Failed to generate JWT:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	log.Println("JWT generated successfully")

	c.JSON(http.StatusOK, gin.H{
		"token": token,
		"user":  u,
	})
}

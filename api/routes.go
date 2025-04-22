package api

import (
	"net/http"

	"github.com/dionisioedu/StickerVerse/internal/auth"
	"github.com/dionisioedu/StickerVerse/internal/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.CORSMiddleware())

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Welcome to StickerVerse 👋"})
	})

	r.POST("/auth/google", auth.GoogleAuthHandler)

	authGroup := r.Group("/")
	authGroup.Use(auth.AuthRequired())
	{
		authGroup.GET("/me", func(c *gin.Context) {
			user := c.GetString("user")
			c.JSON(http.StatusOK, gin.H{"user": user})
		})
	}

	return r
}

package api

import (
	"log"
	"net/http"

	"github.com/dionisioedu/StickerVerse/internal/auth"
	"github.com/dionisioedu/StickerVerse/internal/middleware"
	"github.com/dionisioedu/StickerVerse/internal/user" // importa o package com o repositório
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
			authUser, exists := c.Get("user")
			if !exists {
				log.Printf("GET /me Unauthorized!")
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
				return
			}

			log.Printf("Fetching user data for userID: %s", authUser)
			u, ok := authUser.(*user.User)
			if !ok {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user context"})
				return
			}
			c.JSON(http.StatusOK, gin.H{"user": u})
		})
	}

	return r
}

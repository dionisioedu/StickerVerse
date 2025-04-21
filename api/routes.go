package api

import (
	"net/http"

	"github.com/dionisioedu/StickerVerse/internal/user"
	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	r.GET("/", func(c *gin.Context) {
		u, err := user.GetFirstUser()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, u)
	})

	return r
}

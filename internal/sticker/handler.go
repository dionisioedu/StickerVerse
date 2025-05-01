package sticker

import (
	"net/http"

	"github.com/dionisioedu/StickerVerse/internal/user"
	"github.com/gin-gonic/gin"
)

func ListUserStickersHandler(c *gin.Context) {
	authUser, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	u, ok := authUser.(*user.User)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user context"})
		return
	}

	stickers, err := ListStickersByCreator(u.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch stickers"})
		return
	}

	c.JSON(http.StatusOK, stickers)
}

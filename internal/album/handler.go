package album

import (
	"log"
	"net/http"
	"time"

	"github.com/dionisioedu/StickerVerse/internal/user"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CreateAlbumHandler(c *gin.Context) {
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

	var req struct {
		Name        string `json:"name" binding:"required"`
		Description string `json:"description"`
		IsPrivate   bool   `json:"isPrivate"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	album := Album{
		ID:          uuid.New().String(),
		UserID:      u.ID,
		Name:        req.Name,
		Description: req.Description,
		IsPrivate:   req.IsPrivate,
		CreatedAt:   time.Now(),
	}

	err := CreateAlbum(&album)
	if err != nil {
		log.Printf("CreateAlbum error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create album"})
		return
	}

	c.JSON(http.StatusCreated, album)
}

func ListAlbumsByUserHandler(c *gin.Context) {
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

	albums, err := ListAlbumsByUser(u.ID)
	if err != nil {
		log.Printf("ListAlbumsByUser error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to list albums"})
		return
	}

	if albums == nil {
		albums = []Album{}
	}

	c.JSON(http.StatusOK, albums)
}

func GetAlbumByIDHandler(c *gin.Context) {
	albumID := c.Param("id")

	album, err := GetAlbumByID(albumID)
	if err != nil {
		log.Printf("GetAlbumByID error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch album"})
		return
	}

	if album == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Album not found"})
		return
	}

	c.JSON(http.StatusOK, album)
}

func AddStickerToAlbumHandler(c *gin.Context) {
	albumID := c.Param("id")

	var req struct {
		StickerID string  `json:"stickerId" binding:"required"`
		Position  int     `json:"position" binding:"required"`
		SignedBy  *string `json:"signedBy"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	sticker := AlbumSticker{
		ID:        uuid.New().String(),
		AlbumID:   albumID,
		StickerID: req.StickerID,
		Position:  req.Position,
		SignedBy:  req.SignedBy,
	}

	if req.SignedBy != nil {
		sticker.SignedAt = ptrTime(time.Now())
	}

	err := AddStickerToAlbum(&sticker)
	if err != nil {
		log.Printf("AddStickerToAlbum error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add sticker to album"})
		return
	}

	c.JSON(http.StatusCreated, sticker)
}

func ListStickersInAlbumHandler(c *gin.Context) {
	albumID := c.Param("id")

	stickers, err := ListStickersInAlbum(albumID)
	if err != nil {
		log.Printf("ListStickersInAlbum error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to list stickers"})
		return
	}

	c.JSON(http.StatusOK, stickers)
}

func ptrTime(t time.Time) *time.Time {
	return &t
}

package sticker

import (
	"fmt"

	"github.com/dionisioedu/StickerVerse/internal/db"
)

func ListStickersByCreator(userID string) ([]Sticker, error) {
	var stickers []Sticker
	err := db.DB.Select(&stickers, `
		SELECT * FROM stickers WHERE creator_id = $1 ORDER BY created at DESC
	`, userID)

	if err != nil {
		return nil, fmt.Errorf("ListStickersByCreator error: %w", err)
	}

	return stickers, nil
}

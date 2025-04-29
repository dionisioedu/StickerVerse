package album

import (
	"fmt"

	"github.com/dionisioedu/StickerVerse/internal/db"
)

func CreateAlbum(a *Album) error {
	query := `
        INSERT INTO albums (id, user_id, name, description, is_private, created_at)
        VALUES ($1, $2, $3, $4, $5, $6)
    `
	_, err := db.DB.Exec(query, a.ID, a.UserID, a.Name, a.Description, a.IsPrivate, a.CreatedAt)
	if err != nil {
		return fmt.Errorf("CreateAlbum error: %w", err)
	}
	return nil
}

func GetAlbumByID(id string) (*Album, error) {
	var a Album
	err := db.DB.Get(&a, "SELECT * FROM albums WHERE id = $1", id)
	if err != nil {
		return nil, fmt.Errorf("GetAlbumByID error: %w", err)
	}
	return &a, nil
}

func ListAlbumsByUser(userID string) ([]Album, error) {
	var albums []Album
	err := db.DB.Select(&albums, "SELECT * FROM albums WHERE user_id = $1 ORDER BY created_at DESC", userID)
	if err != nil {
		return nil, fmt.Errorf("ListAlbumsByUser error: %w", err)
	}
	return albums, nil
}

func AddStickerToAlbum(s *AlbumSticker) error {
	query := `
        INSERT INTO album_stickers (id, album_id, sticker_id, position, signed_by, signed_at)
        VALUES ($1, $2, $3, $4, $5, $6)
    `
	_, err := db.DB.Exec(query, s.ID, s.AlbumID, s.StickerID, s.Position, s.SignedBy, s.SignedAt)
	if err != nil {
		return fmt.Errorf("AddStickerToAlbum error: %w", err)
	}
	return nil
}

func ListStickersInAlbum(albumID string) ([]AlbumSticker, error) {
	var stickers []AlbumSticker
	err := db.DB.Select(&stickers, "SELECT * FROM album_stickers WHERE album_id = $1 ORDER BY position ASC", albumID)
	if err != nil {
		return nil, fmt.Errorf("ListStickersInAlbum error: %w", err)
	}
	return stickers, nil
}

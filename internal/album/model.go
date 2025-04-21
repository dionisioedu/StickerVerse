package album

import "time"

type Album struct {
	ID          string    `db:"id" json:"id"`
	UserID      string    `db:"user_id" json:"userId"`
	Name        string    `db:"name" json:"name"`
	Description string    `db:"description" json:"description"`
	IsPrivate   bool      `db:"is_private" json:"isPrivate"`
	CreatedAt   time.Time `db:"created_at" json:"createdAt"`
}

type AlbumSticker struct {
	ID        string     `db:"id" json:"id"`
	AlbumID   string     `db:"album_id" json:"albumId"`
	StickerID string     `db:"sticker_id" json:"stickerId"`
	Position  int        `db:"position" json:"position"`
	SignedBy  *string    `db:"signed_by" json:"signedBy,omitempty"`
	SignedAt  *time.Time `db:"signed_at" json:"signedAt,omitempty"`
}

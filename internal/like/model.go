package like

import "time"

type Like struct {
	ID        string    `db:"id" json:"id"`
	UserID    string    `db:"user_id" json:"userId"`
	StickerID string    `db:"sticker_id" json:"stickerId"`
	CreatedAt time.Time `db:"created_at" json:"createdAt"`
}

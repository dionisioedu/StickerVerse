package sticker

import "time"

type Sticker struct {
	ID        string    `db:"id" json:"id"`
	CreatorID string    `db:"creator_id" json:"creatorId"`
	Title     *string   `db:"title" json:"title"`
	ImageURL  string    `db:"image_url" json:"imageUrl"`
	Rarity    string    `db:"rarity" json:"rarity"` // common, rare, legendary
	CreatedAt time.Time `db:"created_at" json:"createdAt"`
}

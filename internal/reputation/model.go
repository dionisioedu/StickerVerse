package reputation

import "time"

type Reputation struct {
	ID        string    `db:"id" json:"id"`
	FromUser  string    `db:"from_user" json:"fromUser"`
	ToUser    string    `db:"to_user" json:"toUser"`
	Trait     string    `db:"trait" json:"trait"` // legal, sexy, confiavel, criativo, divertido
	CreatedAt time.Time `db:"created_at" json:"createdAt"`
}

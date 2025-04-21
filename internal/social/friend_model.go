package social

import "time"

type Friendship struct {
	ID          string    `db:"id" json:"id"`
	RequesterID string    `db:"requester_id" json:"requesterId"`
	AddresseeID string    `db:"addressee_id" json:"addresseeId"`
	Status      string    `db:"status" json:"status"` // pending, accepted, declined
	CreatedAt   time.Time `db:"created_at" json:"createdAt"`
}

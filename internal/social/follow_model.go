package social

import "time"

type Follow struct {
	ID          string    `db:"id" json:"id"`
	FollowerID  string    `db:"follower_id" json:"followerId"`
	FollowingID string    `db:"following_id" json:"followingId"`
	CreatedAt   time.Time `db:"created_at" json:"createdAt"`
}

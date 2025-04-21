package user

import (
	"time"
)

type User struct {
	ID         string    `db:"id" json:"id"`
	Username   string    `db:"username" json:"username"`
	Password   string    `db:"password_hash" json:"passwordHash"`
	Email      string    `db:"email" json:"email"`
	AvatarURL  *string   `db:"avatar_url" json:"avatarUrl"`
	Bio        *string   `db:"bio" json:"bio"`
	Provider   string    `db:"provider" json:"provider"`
	ProviderID string    `db:"provider_id" json:"providerId"`
	CreatedAt  time.Time `db:"created_at" json:"createdAt"`
	UpdatedAt  time.Time `db:"updated_at" json:"updatedAt"`
}

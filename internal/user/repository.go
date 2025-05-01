package user

import (
	"fmt"
	"time"

	"github.com/dionisioedu/StickerVerse/internal/db"
	"github.com/dionisioedu/StickerVerse/internal/model"
)

func GetFirstUser() (*User, error) {
	var u User
	err := db.DB.Get(&u, "SELECT * FROM users LIMIT 1")
	if err != nil {
		return nil, fmt.Errorf("GetFirstUser error: %w", err)
	}
	return &u, nil
}

func GetUserByID(userID string) (*User, error) {
	var u User
	err := db.DB.Get(&u, "SELECT * FROM users WHERE id = $1", userID)
	if err != nil {
		return nil, fmt.Errorf("GetUserByID error: %w", err)
	}
	return &u, nil
}

func FindOrCreateUser(googleUser *model.GoogleUser) (*User, error) {
	var u User
	err := db.DB.Get(&u, "SELECT * FROM users WHERE provider = 'google' AND provider_id = $1", googleUser.Sub)
	if err == nil {
		return &u, nil
	}

	id := googleUser.Sub
	now := time.Now()

	query := `
        INSERT INTO users (id, username, email, avatar_url, provider, provider_id, created_at, updated_at)
        VALUES ($1, $2, $3, $4, 'google', $5, $6, $6)
        RETURNING *`
	err = db.DB.Get(&u, query,
		id,
		googleUser.Name,
		googleUser.Email,
		googleUser.Picture,
		googleUser.Sub,
		now,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to insert user: %w", err)
	}

	return &u, nil
}

func UpdateUserDisplay(userID string, display string) error {
	_, err := db.DB.Exec(
		`UPDATE users SET display = $1, updated_at = now() WHERE id = $2`,
		display, userID,
	)
	if err != nil {
		return fmt.Errorf("UpdateUserDisplay error: %w", err)
	}
	return nil
}

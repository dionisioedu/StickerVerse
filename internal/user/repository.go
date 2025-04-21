package user

import (
	"fmt"

	"github.com/dionisioedu/StickerVerse/internal/db"
)

func GetFirstUser() (*User, error) {
	var u User
	err := db.DB.Get(&u, "SELECT * FROM users LIMIT 1")
	if err != nil {
		return nil, fmt.Errorf("GetFirstUser error: %w", err)
	}
	return &u, nil
}

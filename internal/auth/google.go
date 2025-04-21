package auth

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/dionisioedu/StickerVerse/internal/model"
)

func VerifyGoogleToken(idToken string) (*model.GoogleUser, error) {
	url := fmt.Sprintf("https://oauth2.googleapis.com/tokeninfo?id_token=%s", idToken)
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode != 200 {
		return nil, fmt.Errorf("invalid token")
	}
	defer resp.Body.Close()

	var gUser model.GoogleUser
	if err := json.NewDecoder(resp.Body).Decode(&gUser); err != nil {
		return nil, err
	}

	return &gUser, nil
}

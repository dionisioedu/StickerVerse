package main

import (
	"log"
	"net/http"
	"os"

	"github.com/dionisioedu/StickerVerse/api"
	"github.com/dionisioedu/StickerVerse/internal/db"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	db.ConnectDB()

	r := api.SetupRoutes()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("StickerVerse running on http://localhost:%s", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}

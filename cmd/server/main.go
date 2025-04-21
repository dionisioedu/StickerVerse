package main

import (
    "log"
    "net/http"
    "os"

    "github.com/joho/godotenv"
    "github.com/seu-usuario/StickerVerse/api"
)

func main() {
    godotenv.Load()

    r := api.SetupRoutes()

    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    log.Printf("StickerVerse running on http://localhost:%s", port)
    log.Fatal(http.ListenAndServe(":"+port, r))
}

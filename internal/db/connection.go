package db

import (
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
)

var DB *sqlx.DB

func ConnectDB() {
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("DATABASE_URL not set")
	}

	var err error
	DB, err = sqlx.Open("pgx", dbURL)
	if err != nil {
		log.Fatalf("Error opening DB: %v", err)
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)
	DB.SetConnMaxLifetime(5 * time.Minute)

	if err = DB.Ping(); err != nil {
		log.Fatalf("Error pinging DB: %v", err)
	}

	fmt.Println("✅ Connected to PostgreSQL!")
}

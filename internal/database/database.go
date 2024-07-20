package database

import (
    "log"
    "os"
    "github.com/jmoiron/sqlx"
    _ "github.com/lib/pq"
)

func InitDB() *sqlx.DB {
    dsn := os.Getenv("DATABASE_URL")
    db, err := sqlx.Connect("postgres", dsn)
    if err != nil {
        log.Fatalf("Failed to connect to database: %v", err)
    }
    return db
}

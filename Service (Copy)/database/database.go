package database

import (
    "database/sql"
    "fmt"
    "log"

    "service/config"
    _ "github.com/lib/pq"
)

// InitDB initializes and returns a DB connection pool
func InitDB(cfg *config.Config) *sql.DB {
    dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
        cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName)

    db, err := sql.Open("postgres", dsn)
    if err != nil {
        log.Fatalf("Failed to open DB: %v", err)
    }

    err = db.Ping()
    if err != nil {
        log.Fatalf("Failed to connect to DB: %v", err)
    }

    log.Println("✅ Connected to PostgreSQL")
    return db
}

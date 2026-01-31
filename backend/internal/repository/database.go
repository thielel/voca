package repository

import (
	"database/sql"
	"log"
	"os"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
)

// InitDB initializes the SQLite database and runs migrations
func InitDB(dbPath string) (*sql.DB, error) {
	// Ensure the directory exists
	dir := filepath.Dir(dbPath)
	if dir != "." && dir != "" {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return nil, err
		}
	}

	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}

	// Test the connection
	if err := db.Ping(); err != nil {
		return nil, err
	}

	// Run migrations
	if err := runMigrations(db); err != nil {
		return nil, err
	}

	log.Printf("Database initialized at %s", dbPath)
	return db, nil
}

// runMigrations creates the necessary tables
func runMigrations(db *sql.DB) error {
	migration := `
		CREATE TABLE IF NOT EXISTS personality_results (
			id TEXT PRIMARY KEY,
			session_id TEXT NOT NULL,
			extraversion REAL NOT NULL,
			agreeableness REAL NOT NULL,
			conscientiousness REAL NOT NULL,
			emotional_stability REAL NOT NULL,
			openness REAL NOT NULL,
			created_at TEXT NOT NULL DEFAULT (datetime('now'))
		);

		CREATE INDEX IF NOT EXISTS idx_personality_results_session_id 
		ON personality_results(session_id);

		CREATE INDEX IF NOT EXISTS idx_personality_results_created_at 
		ON personality_results(created_at DESC);

		CREATE TABLE IF NOT EXISTS trait_interpretations (
			id TEXT PRIMARY KEY,
			result_id TEXT NOT NULL REFERENCES personality_results(id),
			trait TEXT NOT NULL,
			interpretation TEXT NOT NULL,
			created_at TEXT NOT NULL DEFAULT (datetime('now')),
			UNIQUE(result_id, trait)
		);

		CREATE INDEX IF NOT EXISTS idx_trait_interpretations_result_id 
		ON trait_interpretations(result_id);
	`

	_, err := db.Exec(migration)
	if err != nil {
		return err
	}

	log.Println("Database migrations completed")
	return nil
}

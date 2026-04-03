package store

import (
	"context"
	"database/sql"
	"fmt"

	_ "modernc.org/sqlite" // Underscore registers this drive with the database/sql package

	"os"
	"time"
)

type Store struct {
	db *sql.DB // "I have a field called db, which is a sql.DB connection"
}

type Entry struct {
	Title     string
	Input     string
	Output    string
	Model     string
	CreatedAt time.Time
	RevealAt  time.Time
	Visible   bool
}

type EntryRow struct {
	ID int
	Entry
}

func InitDatabase(dbPath string) (*Store, error) {
	db, err := sql.Open("sqlite", dbPath)
	if err != nil {
		return nil, err
	}
	schema, err := os.ReadFile("internal/store/schema.sql")
	if err != nil {
		return nil, err
	}
	_, err = db.ExecContext(context.Background(), string(schema))
	if err != nil {
		return nil, err
	}
	return &Store{db: db}, nil
}

func (store *Store) CreateEntry(entry *Entry) (int64, error) {
	if entry.Input == "" {
		return 0, fmt.Errorf("input cannot be empty")
	}
	revealAt := time.Now().Add(time.Hour)

	result, err := store.db.ExecContext(
		context.Background(),
		`INSERT INTO entries (title, input, output, model, reveal_at) VALUES (?, ?, ?, ?, ?)`,
		entry.Title, entry.Input, entry.Output, entry.Model, revealAt,
	)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

//func update

//func delete

//func llmOutput

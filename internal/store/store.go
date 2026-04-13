package store

import (
	"bufio"
	"context"
	"database/sql"
	"fmt"
	"strings"

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

func (store *Store) Input() (*Entry, error) {
	// Get user inputs and save as strings
	var title, body string
	var builder strings.Builder
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Title:\n")
	// Read a single line
	if scanner.Scan() {
		title = scanner.Text()
	}

	fmt.Print("Body: (type 'exit' to end input)\n")
	// Read multiple lines
	for scanner.Scan() {
		body := scanner.Text()
		if body == "exit" {
			break
		}
		builder.WriteString(body + "\n") // Add to builder
	}
	body = builder.String() // Get complete text with all lines
	if body == "" {
		return nil, fmt.Errorf("Body cannot be empty")
	}

	// Build user inputs into the Entry struct
	entry := &Entry{
		Title: title,
		Input: body,
	}

	fmt.Println("===== RECORD SAVED =====\n", title, "\n", body)

	return entry, nil
}

func (store *Store) AllEntries() ([]Entry, error) {
	var entries []Entry

	rows, err := store.db.QueryContext(
		context.Background(),
		`SELECT id, title, input, output, model, created_at, reveal_at, visible FROM entries;`,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var e EntryRow
		err := rows.Scan(&e.ID, &e.Title, &e.Input, &e.Output, &e.Model, &e.CreatedAt, &e.RevealAt, &e.Visible)
		if err != nil {
			return nil, err
		}
		entries = append(entries, e.Entry)
	}

	return entries, nil
}

func (store *Store) EntryByModel(model string) ([]Entry, error) {
	var entries []Entry

	rows, err := store.db.QueryContext(
		context.Background(),
		`SELECT * FROM entries WHERE model=?;`, model,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var e EntryRow
		err := rows.Scan(&e.ID, &e.Title, &e.Input, &e.Output, &e.Model, &e.CreatedAt, &e.RevealAt, &e.Visible)
		if err != nil {
			return nil, err
		}
		entries = append(entries, e.Entry)
	}

	return entries, nil
}

func (store *Store) View() ([]Entry, error) {
	var entries []Entry

	rows, err := store.db.QueryContext(
		context.Background(),
		`SELECT * FROM entries;`,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var e EntryRow
		err := rows.Scan(&e.ID, &e.Title, &e.Input, &e.Output)
		if err != nil {
			return nil, err
		}
		entries = append(entries, e.Entry)
	}

	return entries, nil
}

//func update

//func delete

//func llmOutput

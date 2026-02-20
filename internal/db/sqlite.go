package db

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"

	_ "modernc.org/sqlite"
)

type Storage struct {
	db *sql.DB
}

func Open(dbPath string) (*Storage, error) {
	// Создаем директорию для файла БД (если нужна)
	if dbPath != ":memory:" {
		dir := filepath.Dir(dbPath)
		if dir != "." && dir != "" {
			if err := os.MkdirAll(dir, 0o755); err != nil {
				return nil, fmt.Errorf("create db dir %s: %w", dir, err)
			}
		}
	}

	var dsn string

	if dbPath == ":memory:" {
		dsn = "file::memory:?_pragma=foreign_keys(1)"
	} else {
		dsn = fmt.Sprintf("file:%s?_pragma=foreign_keys(1)", dbPath)
	}

	database, err := sql.Open("sqlite", dsn)
	if err != nil {
		return nil, fmt.Errorf("open sqlite: %w", err)
	}

	if err := database.Ping(); err != nil {
		_ = database.Close()
		return nil, fmt.Errorf("ping sqlite: %w", err)
	}

	return &Storage{db: database}, nil
}

func (s *Storage) DB() *sql.DB {
	return s.db
}

func (s *Storage) Close() error {
	return s.db.Close()
}
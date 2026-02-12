package db

import (
	"fmt"
	"os"
)

func RunMigration(storage *Storage, path string) error {
	sqlBytes, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("read migration: %w", err)
	}

	if _, err := storage.db.Exec(string(sqlBytes)); err != nil {
		return fmt.Errorf("exec migration: %w", err)
	}

	return nil
}

package repo

import (
	"context"
	"database/sql"
	"errors"
	"financetracker/internal/db"
	"financetracker/internal/models"
	"fmt"
	"strings"

	sqliteDriver "modernc.org/sqlite"
	sqlite3 "modernc.org/sqlite/lib"
)

type CategoryRepo struct {
	db *sql.DB
}

func NewCategoryRepo(storage *db.Storage) *CategoryRepo {
	return &CategoryRepo{db: storage.DB()}
}

func (r *CategoryRepo) Create(ctx context.Context, c *models.Category) (int64, error) {
	const q = `INSERT INTO categories(name, type) VALUES(?, ?)`
	res, err := r.db.ExecContext(ctx, q, c.Name, c.Type)
	if err != nil {
		var sqlErr *sqliteDriver.Error
		if errors.As(err, &sqlErr) &&
			sqlErr.Code()&0xFF == sqlite3.SQLITE_CONSTRAINT &&
			strings.Contains(err.Error(), "UNIQUE constraint failed: categories.name") {
			return 0, ErrCategoryAlreadyExists
		}
		return 0, fmt.Errorf("create category: %w", err)
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("category last insert id: %w", err)
	}
	return id, nil
}

func (r *CategoryRepo) GetByID(ctx context.Context, id int64) (*models.Category, error) {
	const q = `SELECT id, name, type FROM categories WHERE id = ?`

	var c models.Category
	err := r.db.QueryRowContext(ctx, q, id).Scan(&c.ID, &c.Name, &c.Type)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("get category by id: %w", err)
	}
	return &c, nil
}

func (r *CategoryRepo) List(ctx context.Context) ([]models.Category, error) {
	const q = `SELECT id, name, type FROM categories ORDER BY id`

	rows, err := r.db.QueryContext(ctx, q)
	if err != nil {
		return nil, fmt.Errorf("list categories: %w", err)
	}
	defer rows.Close()

	var out []models.Category
	for rows.Next() {
		var c models.Category
		if err := rows.Scan(&c.ID, &c.Name, &c.Type); err != nil {
			return nil, fmt.Errorf("scan category: %w", err)
		}
		out = append(out, c)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows category: %w", err)
	}
	return out, nil
}

func (r *CategoryRepo) Delete(ctx context.Context, id int64) (error) {
	const q = "DELETE FROM categories WHERE id = ?"
	const sqliteConstraintForeignKey = 787 // SQLITE_CONSTRAINT_FOREIGNKEY

	result, err := r.db.ExecContext(ctx, q, id)
	if err != nil {
		var sqlErr *sqliteDriver.Error
		if errors.As(err, &sqlErr) &&
		sqlErr.Code()&0xFF == sqlite3.SQLITE_CONSTRAINT &&
		strings.Contains(err.Error(), "FOREIGN KEY constraint failed") {
			return ErrCategoryInUse
		}
		return fmt.Errorf("delete category: %w", err)
	}

	n, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("delete category rows affected: %w", err)
	}

	if n == 0 {
		return ErrCategoryNotFound
	}

	return nil
}
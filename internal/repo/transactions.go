package repo

import (
	"context"
	"database/sql"
	"financetracker/internal/db"
	"financetracker/internal/models"
	"fmt"
)

type TransactionsRepo struct {
	db *sql.DB
}

func NewTransactionsRepo(storage *db.Storage) *TransactionsRepo {
	return &TransactionsRepo{db: storage.DB()}
}

func (r *TransactionsRepo) Create(ctx context.Context, t *models.Transaction) (int64, error) {
	const q = `INSERT INTO transactions(amount, type, category_id, date, description) VALUES(?, ?, ?, ?, ?)`
	res, err := r.db.ExecContext(ctx, q, t.Amount, t.Type, t.CategoryID, t.Date, t.Description)
	if err != nil {
		return 0, fmt.Errorf("create transaction: %w", err)
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("transaction last insert id: %w", err)
	}
	return id, nil
}

func (r *TransactionsRepo) GetByID(ctx context.Context, id int64) (*models.Transaction, error) {
	const q = `SELECT id, amount, type, category_id, date, description FROM transactions WHERE id = ?`

	var t models.Transaction
	err := r.db.QueryRowContext(ctx, q, id).Scan(&t.ID, &t.Amount, &t.Type, &t.CategoryID, &t.Date, &t.Description)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("get transaction by id: %w", err)
	}
	return &t, nil
}

func (r *TransactionsRepo) List(ctx context.Context) ([]models.Transaction, error) {
	const q = `SELECT id, amount, type, category_id, date, description FROM transactions ORDER BY id`

	rows, err := r.db.QueryContext(ctx, q)
	if err != nil {
		return nil, fmt.Errorf("list transactions: %w", err)
	}
	defer rows.Close()

	var out []models.Transaction
	for rows.Next() {
		var t models.Transaction
		if err := rows.Scan(&t.ID, &t.Amount, &t.Type, &t.CategoryID, &t.Date, &t.Description); err != nil {
			return nil, fmt.Errorf("scan transaction: %w", err)
		}
		out = append(out, t)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows transaction: %w", err)
	}
	return out, nil
}

func (r *TransactionsRepo) Delete(ctx context.Context, id int64) (error) {
	const q = "DELETE FROM transactions WHERE id = ?"

	result, err := r.db.ExecContext(ctx, q, id)
	if err != nil {
		return fmt.Errorf("delete transaction: %w", err)
	}

	n, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("delete transaction rows affected: %w", err)
	}

	if n == 0 {
		return ErrTransactionNotFound
	}

	return nil
}

func (r *TransactionsRepo) Update(ctx context.Context, t *models.Transaction) error {
	if t == nil {
		return fmt.Errorf("update transaction: nil transaction")
	}
	
	const q = "UPDATE transactions SET amount = ?, date = ?, description = ? WHERE id = ?"

	res, err := r.db.ExecContext(ctx, q, t.Amount, t.Date, t.Description, t.ID)
	if err != nil {
		return fmt.Errorf("update transaction: %w", err)
	}

	n, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("update transaction rows affected: %w", err)
	}

	if n == 0 {
		return ErrTransactionNotFound
	}

	return nil
}
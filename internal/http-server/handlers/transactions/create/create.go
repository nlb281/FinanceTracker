package create

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"
	"strings"
	"time"

	"financetracker/internal/models"
)

type TransactionCreator interface {
	Create(ctx context.Context, t *models.Transaction) (int64, error)
}

type CategoryGetterByID interface {
	GetByID(ctx context.Context, id int64) (*models.Category, error)
}

type Response struct {
	Status string `json:"status"`
	Error  string `json:"error,omitempty"`
	ID     int64  `json:"id,omitempty"`
}

func New(creator TransactionCreator, categoryGetterByID CategoryGetterByID) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.transactions.create"

		var input models.Transaction
		if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
			slog.Info("invalid json", "op", op, "error", err)
			writeJSON(w, http.StatusBadRequest, Response{Status: "error", Error: "invalid json"})
			return
		}

		input.Type = strings.ToLower(strings.TrimSpace(input.Type))
		input.Date = strings.TrimSpace(input.Date)
		input.Description = strings.TrimSpace(input.Description)

		if input.Amount < 0 {
			writeJSON(w, http.StatusBadRequest, Response{Status: "error", Error: "amount must be >= 0"})
			return
		}
		if input.Type != "income" && input.Type != "expense" {
			writeJSON(w, http.StatusBadRequest, Response{Status: "error", Error: "type must be income or expense"})
			return
		}
		if input.CategoryID <= 0 {
			writeJSON(w, http.StatusBadRequest, Response{Status: "error", Error: "category_id must be > 0"})
			return
		}
		if _, err := time.Parse("2006-01-02", input.Date); err != nil {
			writeJSON(w, http.StatusBadRequest, Response{Status: "error", Error: "date must be YYYY-MM-DD"})
			return
		}

		cat, err := categoryGetterByID.GetByID(r.Context(), input.CategoryID)
		if err != nil {
			slog.Info("failed to get category", "op", op, "category_id", input.CategoryID, "error", err)
			writeJSON(w, http.StatusInternalServerError, Response{Status: "error", Error: "failed to get category"})
			return
		}
		if cat == nil {
			writeJSON(w, http.StatusNotFound, Response{Status: "error", Error: "category not found"})
			return
		}
		if cat.Type != input.Type {
			writeJSON(w, http.StatusBadRequest, Response{Status: "error", Error: "transaction type must match category type"})
			return
		}

		id, err := creator.Create(r.Context(), &input)
		if err != nil {
			slog.Info("create transaction failed", "op", op, "error", err)
			writeJSON(w, http.StatusInternalServerError, Response{Status: "error", Error: "failed to create transaction"})
			return
		}

		writeJSON(w, http.StatusCreated, Response{Status: "ok", ID: id})
	}
}

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}

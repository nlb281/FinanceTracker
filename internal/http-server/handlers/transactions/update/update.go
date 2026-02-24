package update

import (
	"context"
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"
	"strconv"
	"strings"
	"time"

	"financetracker/internal/models"
	"financetracker/internal/repo"

	"github.com/go-chi/chi/v5"
)

type TransactionUpdater interface {
	Update(ctx context.Context, t *models.Transaction) (error)
}

type TransactionGetterByID interface {
	GetByID(ctx context.Context, id int64) (*models.Transaction, error)
}

type Response struct {
	Status string `json:"status"`
	Error  string `json:"error,omitempty"`
}

type patchInput struct {
	Amount      *float64 `json:"amount"`
	Date        *string  `json:"date"`
	Description *string  `json:"description"`
}

func New(updater TransactionUpdater, getter TransactionGetterByID) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.transactions.update"

		idStr := chi.URLParam(r, "id")
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil || id <= 0 {
			writeJSON(w, http.StatusBadRequest, Response{Status: "error", Error: "invalid id"})
			return
		}

		var input patchInput
		if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
			slog.Info("invalid json", "op", op, "error", err)
			writeJSON(w, http.StatusBadRequest, Response{Status: "error", Error: "invalid json"})
			return
		}

		current, err := getter.GetByID(r.Context(), id)
		if err != nil {
			slog.Info("get transaction failed", "op", op, "id", id, "error", err)
			writeJSON(w, http.StatusInternalServerError, Response{Status: "error", Error: "internal error"})
			return
		}
		if current == nil {
			writeJSON(w, http.StatusNotFound, Response{Status: "error", Error: "transaction not found"})
			return
		}

		updated := *current
		changed := false

		if input.Amount != nil {
			if *input.Amount < 0 {
				writeJSON(w, http.StatusBadRequest, Response{Status: "error", Error: "amount must be >= 0"})
				return
			}
			updated.Amount = *input.Amount
			changed = true
		}

		if input.Date != nil {
			date := strings.TrimSpace(*input.Date)
			if _, err := time.Parse("2006-01-02", date); err != nil {
				writeJSON(w, http.StatusBadRequest, Response{Status: "error", Error: "date must be YYYY-MM-DD"})
				return
			}
			updated.Date = date
			changed = true
		}

		if input.Description != nil {
			updated.Description = strings.TrimSpace(*input.Description)
			changed = true
		}

		if !changed {
			writeJSON(w, http.StatusBadRequest, Response{Status: "error", Error: "nothing to update"})
			return
		}

		err = updater.Update(r.Context(), &updated)
		if err != nil {
			if errors.Is(err, repo.ErrTransactionNotFound) {
				writeJSON(w, http.StatusNotFound, Response{Status: "error", Error: "transaction not found"})
				return
			}
			slog.Info("update transaction failed", "op", op, "id", id, "error", err)
			writeJSON(w, http.StatusInternalServerError, Response{Status: "error", Error: "internal error"})
			return
		}

		writeJSON(w, http.StatusOK, Response{Status: "ok"})
	}
}

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}

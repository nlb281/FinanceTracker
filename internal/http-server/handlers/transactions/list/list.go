package list

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"

	"financetracker/internal/models"
)

type TransactionLister interface {
	List(ctx context.Context) ([]models.Transaction, error)
}

type Response struct {
	Status string `json:"status"`
	Error  string `json:"error,omitempty"`
}

func New(lister TransactionLister) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.transactions.list"

		items, err := lister.List(r.Context())
		if err != nil {
			slog.Info("list transactions failed", "op", op, "error", err)
			writeJSON(w, http.StatusInternalServerError, Response{Status: "error", Error: "failed to list transactions"})
			return
		}

		writeJSON(w, http.StatusOK, items)
	}
}

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}

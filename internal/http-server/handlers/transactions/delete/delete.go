package delete

import (
	"context"
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"
	"strconv"

	"financetracker/internal/repo"

	"github.com/go-chi/chi/v5"
)

type TransactionDeleter interface {
	Delete(ctx context.Context, id int64) error
}

type Response struct {
	Status string `json:"status"`
	Error  string `json:"error,omitempty"`
	ID     int64  `json:"id,omitempty"`
}

func New(deleter TransactionDeleter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.transactions.delete"

		idStr := chi.URLParam(r, "id")
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil || id <= 0 {
			writeJSON(w, http.StatusBadRequest, Response{Status: "error", Error: "invalid id"})
			return
		}

		err = deleter.Delete(r.Context(), id)
		if err != nil {
			if errors.Is(err, repo.ErrTransactionNotFound) {
				writeJSON(w, http.StatusNotFound, Response{Status: "error", Error: "transaction not found"})
				return
			}
			slog.Info("delete transaction failed", "op", op, "id", id, "error", err)
			writeJSON(w, http.StatusInternalServerError, Response{Status: "error", Error: "internal error"})
			return
		}

		writeJSON(w, http.StatusOK, Response{Status: "ok", ID: id})
	}
}

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}

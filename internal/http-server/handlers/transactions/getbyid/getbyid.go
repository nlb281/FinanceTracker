package getbyid

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"
	"strconv"

	"financetracker/internal/models"

	"github.com/go-chi/chi/v5"
)

type TransactionGetterByID interface {
	GetByID(ctx context.Context, id int64) (*models.Transaction, error)
}

type Response struct {
	Status string `json:"status"`
	Error  string `json:"error,omitempty"`
}

func New(getterByID TransactionGetterByID) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.transactions.getbyid"

		idStr := chi.URLParam(r, "id")
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil || id <= 0 {
			writeJSON(w, http.StatusBadRequest, Response{Status: "error", Error: "invalid id"})
			return
		}

		item, err := getterByID.GetByID(r.Context(), id)
		if err != nil {
			slog.Info("get transaction by id failed", "op", op, "id", id, "error", err)
			writeJSON(w, http.StatusInternalServerError, Response{Status: "error", Error: "failed to get transaction"})
			return
		}
		if item == nil {
			writeJSON(w, http.StatusNotFound, Response{Status: "error", Error: "transaction not found"})
			return
		}

		writeJSON(w, http.StatusOK, item)
	}
}

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}

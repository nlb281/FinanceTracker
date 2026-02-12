package list

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"

	"financetracker/internal/models"
)

type CategoryLister interface {
	List(ctx context.Context) ([]models.Category, error)
}

type Response struct {
	Status string `json:"status"`
	Error  string `json:"error,omitempty"`
	ID     int64  `json:"id,omitempty"`
}

func New(lister CategoryLister) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.categories.list"

		items, err := lister.List(r.Context())
		if err != nil {
			slog.Info("list categories failed", "op", op, "error", err)
			writeJSON(w, http.StatusInternalServerError, Response{Status: "error", Error: "failed to list categories"})
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

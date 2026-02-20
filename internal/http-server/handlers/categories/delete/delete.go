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

type CategoryDeleter interface {
	Delete(ctx context.Context, id int64) error
}

type Response struct {
	Status string `json:"status"`
	Error  string `json:"error,omitempty"`
	ID     int64  `json:"id,omitempty"`
}

func New(deleter CategoryDeleter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.categories.delete"

		idStr := chi.URLParam(r, "id")
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil || id <= 0 {
			slog.Info("invalid category id", "op", op, "id", idStr)
			writeJSON(w, http.StatusBadRequest, Response{Status: "error", Error: "invalid id"})
			return
		}

		err = deleter.Delete(r.Context(), id)
		if err != nil {
			if errors.Is(err, repo.ErrCategoryNotFound) {
				slog.Info("category not found", "op", op, "id", id)
				writeJSON(w, http.StatusNotFound, Response{Status: "error", Error: "category not found"})
				return
			}
			if errors.Is(err, repo.ErrCategoryInUse) {
				slog.Info("category is used in transactions", "op", op, "id", id)
				writeJSON(w, http.StatusConflict, Response{Status: "error", Error: "category is used in transactions"})
				return
			}

			slog.Info("failed to delete category", "op", op, "id", id, "error", err)
			writeJSON(w, http.StatusInternalServerError, Response{Status: "error", Error: "internal error"})
			return
		}

		slog.Info("category deleted", "op", op, "id", id)
		writeJSON(w, http.StatusOK, Response{Status: "ok", ID: id})
	}
}

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}

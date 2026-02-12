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

type CategoryGetterByID interface {
	GetByID(ctx context.Context, id int64) (*models.Category, error)
}

type Response struct {
	Status string `json:"status"`
	Error  string `json:"error,omitempty"`
	ID     int64  `json:"id,omitempty"`
}

func New(getterByID CategoryGetterByID) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.categories.getbyid"

		idStr := chi.URLParam(r, "id")
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil || id <= 0 {
			slog.Info("invalid category id", "op", op, "id", idStr)
			writeJSON(w, http.StatusBadRequest, Response{Status: "error", Error: "invalid id"})
			return
		}

		category, err := getterByID.GetByID(r.Context(), id)
		if err != nil {
			slog.Info("get category by id failed", "op", op, "id", id, "error", err)
			writeJSON(w, http.StatusInternalServerError, Response{Status: "error", Error: "failed to get category"})
			return
		}

		if category == nil {
			slog.Info("category not found", "op", op, "id", id)
			writeJSON(w, http.StatusNotFound, Response{Status: "error", Error: "category not found"})
			return
		}

		slog.Info("category returned", "op", op, "id", id)
		writeJSON(w, http.StatusOK, category)
	}
}

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}

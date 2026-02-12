package create

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"
	"strings"

	"financetracker/internal/models"
)

type CategoryCreator interface {
	Create(ctx context.Context, c *models.Category) (int64, error)
}

type Response struct {
	Status string `json:"status"`
	Error  string `json:"error,omitempty"`
	ID     int64  `json:"id,omitempty"`
}

func New(creator CategoryCreator) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.categories.create"

		var input models.Category
		if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
			slog.Info("invalid json", "op", op, "error", err)
			writeJSON(w, http.StatusBadRequest, Response{Status: "error", Error: "invalid json"})
			return
		}

		input.Name = strings.TrimSpace(input.Name)
		input.Type = strings.ToLower(strings.TrimSpace(input.Type))

		if input.Name == "" {
			slog.Info("name is required", "op", op)
			writeJSON(w, http.StatusBadRequest, Response{Status: "error", Error: "name is required"})
			return
		}
		if input.Type != "income" && input.Type != "expense" {
			slog.Info("invalid category type", "op", op, "type", input.Type)
			writeJSON(w, http.StatusBadRequest, Response{Status: "error", Error: "type must be income or expense"})
			return
		}

		id, err := creator.Create(r.Context(), &input)
		if err != nil {
			slog.Info("create category failed", "op", op, "error", err)
			writeJSON(w, http.StatusInternalServerError, Response{Status: "error", Error: "failed to create category"})
			return
		}

		slog.Info("category created", "op", op, "id", id)
		writeJSON(w, http.StatusCreated, Response{Status: "ok", ID: id})
	}
}


func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}

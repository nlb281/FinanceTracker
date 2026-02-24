package main

import (
	"errors"
	"financetracker/internal/config"
	"financetracker/internal/db"
	"financetracker/internal/repo"
	"fmt"
	"log/slog"
	"net/http"
	"os"

	catcreate "financetracker/internal/http-server/handlers/categories/create"
	catdelete "financetracker/internal/http-server/handlers/categories/delete"
	catgetbyid "financetracker/internal/http-server/handlers/categories/getbyid"
	catlist "financetracker/internal/http-server/handlers/categories/list"

	trxcreate "financetracker/internal/http-server/handlers/transactions/create"
	trxdelete "financetracker/internal/http-server/handlers/transactions/delete"
	trxgetbyid "financetracker/internal/http-server/handlers/transactions/getbyid"
	trxlist "financetracker/internal/http-server/handlers/transactions/list"
	trxupdate "financetracker/internal/http-server/handlers/transactions/update"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	if err := run(); err != nil {
		slog.Error(
			"app failed",
			"error", err,
		)
		os.Exit(1)
	}

	
}

func run() error {
	cfg := config.MustLoad("config/local.yaml")

	storage, err := db.Open(cfg.StoragePath)
	if err != nil {
		return fmt.Errorf("open db finance.db: %w", err)
	}
	defer storage.Close()

	if err := db.RunMigration(storage, cfg.MigrationPath); err != nil {
		return fmt.Errorf("run migration %s: %w", cfg.MigrationPath, err)
	}

	categoryRepo := repo.NewCategoryRepo(storage)
	transactionRepo := repo.NewTransactionsRepo(storage)

	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	// categories
	router.Post("/categories", catcreate.New(categoryRepo))
	router.Get("/categories", catlist.New(categoryRepo))
	router.Get("/categories/{id}", catgetbyid.New(categoryRepo))
	router.Delete("/categories/{id}", catdelete.New(categoryRepo))

	// transactions
	router.Post("/transactions", trxcreate.New(transactionRepo, categoryRepo))
	router.Get("/transactions", trxlist.New(transactionRepo))
	router.Get("/transactions/{id}", trxgetbyid.New(transactionRepo))
	router.Delete("/transactions/{id}", trxdelete.New(transactionRepo))
	// cmd/api/main.go (роут)
router.Patch("/transactions/{id}", trxupdate.New(transactionRepo, transactionRepo))



	

	srv := &http.Server{
		Addr:    cfg.HTTPServer.Address,
		Handler: router,
		ReadTimeout:  cfg.HTTPServer.Timeout,
		WriteTimeout: cfg.HTTPServer.Timeout,
		IdleTimeout:  cfg.HTTPServer.IdleTimeout,
	}

	slog.Info("Server started", "address", srv.Addr)

	if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		return fmt.Errorf("listen and serve: %w", err)
	}

	return nil
}
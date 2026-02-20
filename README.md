# FinanceTracker Backend

Backend-часть приложения для учета финансов: обрабатывает операции с категориями и транзакциями, хранит данные в SQLite и применяет миграции при запуске.

## Стек

- Go
- Chi Router
- SQLite (modernc.org/sqlite)
- cleanenv (конфиг из YAML)

## Запуск проекта

go mod tidy
go run ./cmd/api/main.go

## Конфиг

Используется файл `config/local.yaml`:

- `storage_path`: путь к SQLite БД (по умолчанию `storage/finance.db`)
- `migration_path`: путь к SQL-миграции (по умолчанию `migrations/002_init_v2.sql`)
- `http_server.address`: адрес сервера (по умолчанию `localhost:8082`)

Важно: frontend ожидает backend на `http://localhost:8082` (через `/api` proxy).

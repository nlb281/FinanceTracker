package models

type Category struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"` // income | expense
}

type Transaction struct {
	ID          int64   `json:"id"`
	Amount      float64 `json:"amount"`
	Type        string  `json:"type"` // income | expense
	CategoryID  int64   `json:"category_id"`
	Date        string  `json:"date"` // YYYY-MM-DD
	Description string  `json:"description"`
}

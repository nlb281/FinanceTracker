package repo

import "errors"

var ErrCategoryNotFound = errors.New("category not found")
var ErrTransactionNotFound = errors.New("transaction not found")
var ErrCategoryInUse = errors.New("category is used in transactions")
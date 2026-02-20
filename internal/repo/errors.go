package repo

import "errors"

var ErrCategoryNotFound = errors.New("category not found")
var ErrCategoryInUse = errors.New("category is used in transactions")
var ErrCategoryAlreadyExists = errors.New("category already exists")
var ErrTransactionNotFound = errors.New("transaction not found")
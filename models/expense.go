package models

type Expense struct {
	ExpenseName string
	Value       int64 `validate:"gte=1"`
	Tags        map[string]string
}

package main

import (
	"time"
)

type Expense struct {
	Id          int       `json:"id"`
	Date        time.Time `json:"date"`
	Description string    `json:"description"`
	Amount      int       `json:"amount"`
}

type ExpenseList struct {
	Expenses []Expense `json:"expenses"`
}

func (self ExpenseList) get(idx int) Expense {
	return self.Expenses[idx]
}

func (self ExpenseList) size() int {
	return len(self.Expenses)
}

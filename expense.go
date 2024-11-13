package main

import (
	"fmt"
	"time"
)

type Expense struct {
	Id          int       `json:"id"`
	Date        time.Time `json:"date"`
	Description string    `json:"description"`
	Amount      int       `json:"amount"`
}

func (self Expense) print() {
	fmt.Printf("Id: %d\n", self.Id)
	fmt.Printf("Date: %s\n", self.Date)
	fmt.Printf("Description: %s\n", self.Description)
	fmt.Printf("Amount: %d\n", self.Amount)
}

type ExpenseManager struct {
	Expenses []Expense `json:"expenses"`
}

func (self *ExpenseManager) get(idx int) Expense {
	return self.Expenses[idx]
}

func (self *ExpenseManager) size() int {
	return len(self.Expenses)
}

func (self *ExpenseManager) add(expense Expense) {
	self.Expenses = append(self.Expenses, expense)
}

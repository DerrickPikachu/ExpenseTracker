package main

import (
	"fmt"
	"time"
)

type Expense struct {
	Date        time.Time `json:"date"`
	Description string    `json:"description"`
	Amount      int       `json:"amount"`
}

func (self Expense) print() {
	fmt.Printf("Date: %s\n", self.Date)
	fmt.Printf("Description: %s\n", self.Description)
	fmt.Printf("Amount: %d\n", self.Amount)
}

type ExpenseManager struct {
	Expenses map[int]Expense `json:"expenses"`
	LastId   int             `json:"last_id"`
}

func (self *ExpenseManager) get(idx int) Expense {
	return self.Expenses[idx]
}

func (self *ExpenseManager) getAllId() []int {
	var all_id []int = make([]int, 0)
	for k, _ := range self.Expenses {
		all_id = append(all_id, k)
	}
	return all_id
}

func (self *ExpenseManager) size() int {
	return len(self.Expenses)
}

func (self *ExpenseManager) add(expense Expense) {
	self.LastId += 1
	self.Expenses[self.LastId] = expense
}

func (self *ExpenseManager) delete(idx int) {
	delete(self.Expenses, idx)
}

// func (self *ExpenseManager) isExist(idx int) bool {
// 	if (self.Expenses)
// }

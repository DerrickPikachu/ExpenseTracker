package main

import (
	"fmt"
	"os"
	"time"
)

type ExpenseType int

const (
	Food ExpenseType = iota
	Transport
	Entertainment
	Other
)

var expenseType = map[ExpenseType]string{
	Food:          "Food",
	Transport:     "Transport",
	Entertainment: "Entertainment",
	Other:         "Other",
}

func NewExpenseType(name string) (ExpenseType, error) {
	for key, value := range expenseType {
		if value == name {
			return key, nil
		}
	}
	return Other, fmt.Errorf("invalid expense type: %s", name)
}

func (self ExpenseType) String() string {
	return expenseType[self]
}

type Expense struct {
	Date        time.Time   `json:"date"`
	Description string      `json:"description"`
	Amount      int         `json:"amount"`
	Type        ExpenseType `json:"type"`
}

func (self Expense) print() {
	fmt.Printf("Date: %s\n", self.Date)
	fmt.Printf("Description: %s\n", self.Description)
	fmt.Printf("Amount: %d\n", self.Amount)
	fmt.Printf("Type: %s\n", self.Type)
}

type ExpenseManager struct {
	Expenses map[int]Expense `json:"expenses"`
	LastId   int             `json:"last_id"`
}

func (self *ExpenseManager) get(idx int) Expense {
	if expense, ok := self.Expenses[idx]; ok {
		return expense
	}
	fmt.Printf("Non-exist id\n")
	os.Exit(1)
	return Expense{} // no need but make compiler happy
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

func (self *ExpenseManager) update(idx int, expense Expense) {
	self.Expenses[idx] = expense
}

func (self *ExpenseManager) total() int {
	total := 0
	for _, expense := range self.Expenses {
		total += expense.Amount
	}
	return total
}

func (self *ExpenseManager) month_total(month int) int {
	total := 0
	for _, expense := range self.Expenses {
		if expense.Date.Year() != time.Now().Year() ||
			expense.Date.Month() != time.Month(month) {
			continue
		}
		total += expense.Amount
	}
	return total
}

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

var storage string = "expenses.json"

func write_json(expense_manager *ExpenseManager) {
	json_str, _ := json.Marshal(expense_manager)
	os.WriteFile(storage, json_str, 0644)
}

func main() {
	expense_list := ExpenseManager{}
	expense_list.Expenses = make([]Expense, 0)
	if len(os.Args) == 1 {
		fmt.Printf("Missing specific function\n")
		os.Exit(1)
	} else if os.Args[1] == "add" {
		cmd_args := os.Args[2:]
		expense := Expense{}
		for i := 0; i < len(cmd_args); i += 1 {
			if cmd_args[i][:2] != "--" {
				fmt.Printf("Wrong parameter for function \"add\"\n")
				os.Exit(1)
			}
			if cmd_args[i][2:] == "description" {
				expense.Description = cmd_args[i+1]
				i += 1
			} else if cmd_args[i][2:] == "amount" {
				amount, err := strconv.ParseInt(cmd_args[i+1], 10, 32)
				if err != nil {
					log.Fatal(err)
				}
				expense.Amount = int(amount)
				i += 1
			} else {
				fmt.Printf("Wrong parameter for function \"add\"\n")
				os.Exit(1)
			}
		}
		expense.Id = 0
		expense.Date = time.Now()
		expense.print()
		expense_list.add(expense)
	}

	write_json(&expense_list)
}

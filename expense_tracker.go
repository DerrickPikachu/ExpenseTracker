package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"time"
)

var storage string = "expenses.json"

func read_json() *ExpenseManager {
	var expense_manager *ExpenseManager = new(ExpenseManager)
	expense_manager.Expenses = make(map[int]Expense)
	expense_manager.LastId = 0
	json_file, err := os.Open(storage)
	if err == nil {
		bytes, _ := io.ReadAll(json_file)
		json.Unmarshal(bytes, expense_manager)
	}
	json_file.Close()
	return expense_manager
}

func write_json(expense_manager *ExpenseManager) {
	json_str, _ := json.Marshal(expense_manager)
	os.WriteFile(storage, json_str, 0644)
}

func main() {
	expense_list := read_json()
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
		expense.Date = time.Now()
		expense.print()
		expense_list.add(expense)
	} else if os.Args[1] == "list" {
		for i := 0; i < expense_list.size(); i += 1 {
			fmt.Printf("-----------------------------------------------\n")
			fmt.Printf("ID: %d\n", i+1)
			expense_list.get(i + 1).print()
		}
	}

	write_json(expense_list)
}

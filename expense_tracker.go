package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
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
		id_list := expense_list.getAllId()
		sort.Ints(id_list)
		for _, idx := range id_list {
			fmt.Printf("-----------------------------------------------\n")
			fmt.Printf("ID: %d\n", idx)
			expense_list.get(idx).print()
		}
	} else if os.Args[1] == "delete" {
		cmd_args := os.Args[2:]
		for i := 0; i < len(cmd_args); i += 1 {
			if cmd_args[i][:2] != "--" {
				fmt.Printf("Wrong parameter for functiono \"delete\"\n")
				os.Exit(1)
			}
			if cmd_args[i][2:] == "id" {
				idx, err := strconv.ParseInt(cmd_args[i+1], 10, 32)
				if err != nil {
					log.Fatal(err)
				}
				expense_list.delete(int(idx))
				i += 1
			} else {
				fmt.Printf("Wrong parameter for function \"delete\"")
				os.Exit(1)
			}
		}
	} else if os.Args[1] == "update" {
		cmd_args := os.Args[2:]
		var id int = -1
		var description string = ""
		var amount int = -1
		for i := 0; i < len(cmd_args); i += 1 {
			if cmd_args[i][:2] != "--" {
				fmt.Printf("Wrong parameter for function \"update\"\n")
				os.Exit(1)
			}
			if cmd_args[i][2:] == "id" {
				raw_id, err := strconv.ParseInt(cmd_args[i+1], 10, 32)
				if err != nil {
					log.Fatal(err)
				}
				id = int(raw_id)
				i += 1
			} else if cmd_args[i][2:] == "description" {
				description = cmd_args[i+1]
				i += 1
			} else if cmd_args[i][2:] == "amount" {
				raw_num, err := strconv.ParseInt(cmd_args[i+1], 10, 32)
				if err != nil {
					log.Fatal(err)
				}
				amount = int(raw_num)
				i += 1
			}
		}
		if id == -1 {
			fmt.Printf("Missing id to process update function\n")
			os.Exit(1)
		}
		expense := expense_list.get(id)
		if description != "" {
			expense.Description = description
		}
		if amount != -1 {
			expense.Amount = amount
		}
		expense_list.update(id, expense)
	}

	write_json(expense_list)
}

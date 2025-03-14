package main

import (
	"encoding/csv"
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

func add_expense(cmd_args []string, expense_list *ExpenseManager) {
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
		} else if cmd_args[i][2:] == "category" {
			category := cmd_args[i+1]
			expense_type, err := NewExpenseType(category)
			if err != nil {
				log.Fatal(err)
			}
			expense.Type = expense_type
			i += 1
		} else {
			fmt.Printf("Wrong parameter for function \"add\"\n")
			os.Exit(1)
		}
	}
	expense.Date = time.Now()
	expense.print()
	expense_list.add(expense)
}

func list_expenses(cmd_args []string, expense_list *ExpenseManager) {
	category := ""
	for i := 0; i < len(cmd_args); i += 1 {
		if cmd_args[i][:2] != "--" {
			fmt.Printf("Wrong parameter for function \"list\"\n")
			os.Exit(1)
		}
		if cmd_args[i][2:] == "category" {
			category = cmd_args[i+1]
			i += 1
		} else {
			fmt.Printf("Wrong parameter for function \"list\"\n")
			os.Exit(1)
		}
	}

	id_list := expense_list.getAllId()
	sort.Ints(id_list)
	for _, idx := range id_list {
		if category != "" && expense_list.get(idx).Type.String() != category {
			continue
		}
		fmt.Printf("-----------------------------------------------\n")
		fmt.Printf("ID: %d\n", idx)
		expense_list.get(idx).print()
	}
}

func delete_expense(cmd_args []string, expense_list *ExpenseManager) {
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
}

func update_expense(cmd_args []string, expense_list *ExpenseManager) {
	var id int = -1
	var description string = ""
	var amount int = -1
	var category string = ""
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
		} else if cmd_args[i][2:] == "category" {
			category = cmd_args[i+1]
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
	if category != "" {
		expense_type, err := NewExpenseType(category)
		if err != nil {
			log.Fatal(err)
		}
		expense.Type = expense_type
	}
	expense_list.update(id, expense)
}

func summary(cmd_args []string, expense_list *ExpenseManager) {
	month := -1
	for i := 0; i < len(cmd_args); i += 1 {
		if cmd_args[i][:2] != "--" {
			fmt.Printf("Wrong parameter for function \"summary\"\n")
			os.Exit(1)
		}
		if cmd_args[i][2:] == "month" {
			raw_month, err := strconv.ParseInt(cmd_args[i+1], 10, 32)
			if err != nil {
				log.Fatal(err)
			}
			month = int(raw_month)
			i += 1
		} else {
			fmt.Printf("Total expenses: $%d\n", expense_list.total())
		}
	}

	if month != -1 {
		total := expense_list.month_total(month)
		fmt.Printf("Total expenses for month %s: $%d\n", time.Month(month), total)
	} else {
		fmt.Printf("Total expenses: $%d\n", expense_list.total())
	}
}

func output_csv(cmd_args []string, expense_list *ExpenseManager) {
	file, err := os.Create("expenses.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	data := expense_list.toCsv()
	for _, value := range data {
		if err := writer.Write(value); err != nil {
			fmt.Printf("Error writing data to csv: %v\n", err)
		}
	}
}

func main() {
	expense_list := read_json()
	if len(os.Args) == 1 {
		fmt.Printf("Missing specific function\n")
		os.Exit(1)
	} else if os.Args[1] == "add" {
		add_expense(os.Args[2:], expense_list)
	} else if os.Args[1] == "list" {
		list_expenses(os.Args[2:], expense_list)
	} else if os.Args[1] == "delete" {
		delete_expense(os.Args[2:], expense_list)
	} else if os.Args[1] == "update" {
		update_expense(os.Args[2:], expense_list)
	} else if os.Args[1] == "summary" {
		summary(os.Args[2:], expense_list)
	} else if os.Args[1] == "export" {
		output_csv(os.Args[2:], expense_list)
	} else {
		fmt.Printf("Unknown function\n")
		os.Exit(1)
	}

	write_json(expense_list)
}

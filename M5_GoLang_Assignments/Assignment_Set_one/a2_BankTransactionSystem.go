/*
Exercise 2: Bank Transaction System
Topics Covered: Go Constants, Go Loops, Go Break and Continue, Go Functions, Go
Strings, Go Errors
Case Study:
You need to simulate a bank transaction system with the following features:
1. Account Management: Each account has an ID, name, and balance. Store the
accounts in a slice.
2. Deposit Function: A function to deposit money into an account. Validate if the
deposit amount is greater than zero.
3. Withdraw Function: A function to withdraw money from an account. Ensure the
account has a sufficient balance before proceeding. Return appropriate errors
for invalid amounts or insufficient balance.
4. Transaction History: Maintain a transaction history for each account as a string
slice. Use a loop to display the transaction history when requested.
5. Menu System: Implement a menu-driven program where users can choose
actions like deposit, withdraw, view balance, or exit. Use constants for menu
options and break the loop to exit.
*/

package main

import (
	"errors"
	"fmt"
)

type Account struct {
	ID                 int
	Name               string
	Balance            float64
	TransactionHistory []string
}

var accounts []Account

const (
	CreateAccount      = 1
	Deposit            = 2
	Withdraw           = 3
	ViewBalance        = 4
	TransactionHistory = 5
	Exit               = 6
)

func findAccountByID(id int) (*Account, error) {
	for i := range accounts {
		if accounts[i].ID == id {
			return &accounts[i], nil
		}
	}
	return nil, errors.New("account not found")
}

func createAccount(id int, name string, initialBalance float64) error {
	if id <= 0 {
		return errors.New("account ID must be positive")
	}
	if initialBalance < 0 {
		return errors.New("initial balance cannot be negative")
	}
	for _, acc := range accounts {
		if acc.ID == id {
			return errors.New("account ID must be unique")
		}
	}
	newAccount := Account{
		ID:                 id,
		Name:               name,
		Balance:            initialBalance,
		TransactionHistory: []string{fmt.Sprintf("Account Created with Initial Balance: Rs.%.2f", initialBalance)},
	}
	accounts = append(accounts, newAccount)
	return nil
}

func deposit(id int, amount float64) error {
	if amount <= 0 {
		return errors.New("Amount deposited must be greater than zero")
	}
	account, err := findAccountByID(id)
	if err != nil {
		return err
	}
	account.Balance += amount
	fmt.Printf("Amount Deposited: Rs.%.2f..New Balance: Rs.%.2f\n", amount, account.Balance)
	account.TransactionHistory = append(account.TransactionHistory, fmt.Sprintf("Amount Deposited: Rs.%.2f..New Balance: Rs.%.2f", amount, account.Balance))
	return nil
}

func withdraw(id int, amount float64) error {
	if amount <= 0 {
		return errors.New("Amount withdrawn must be greater than zero")
	}
	account, err := findAccountByID(id)
	if err != nil {
		return err
	}
	if account.Balance < amount {
		return errors.New("insufficient balance")
	}
	account.Balance -= amount
	fmt.Printf("Amount Withdrew: Rs.%.2f..New Balance: Rs.%.2f\n", amount, account.Balance)
	account.TransactionHistory = append(account.TransactionHistory, fmt.Sprintf("Amount Withdrew: Rs.%.2f..New Balance: Rs.%.2f", amount))
	return nil
}

func viewBalance(id int) (float64, error) {
	account, err := findAccountByID(id)
	if err != nil {
		return 0, err
	}
	return account.Balance, nil
}

func viewTransactionHistory(id int) ([]string, error) {
	account, err := findAccountByID(id)
	if err != nil {
		return nil, err
	}
	return account.TransactionHistory, nil
}

func main() {
	accounts = []Account{
		{ID: 1, Name: "Anubhav", Balance: 5000, TransactionHistory: []string{"Account Created with Initial Balance: Rs.5000"}},
		{ID: 2, Name: "Sundeep", Balance: 10000, TransactionHistory: []string{"Account Created with Initial Balance: Rs.10000"}},
	}

	for {
		fmt.Println("\nBank Transaction System")
		fmt.Println("1. Create Account")
		fmt.Println("2. Deposit")
		fmt.Println("3. Withdraw")
		fmt.Println("4. View Balance")
		fmt.Println("5. View Transaction History")
		fmt.Println("6. Exit")
		fmt.Print("Choose an option: ")

		var choice int
		_, err := fmt.Scan(&choice)
		if err != nil {
			fmt.Println("Invalid input. Please enter a valid option.")
			continue
		}

		switch choice {
		case CreateAccount:
			fmt.Print("Enter New Account ID: ")
			var id int
			fmt.Scan(&id)
			fmt.Print("Enter Account Holder Name: ")
			var name string
			fmt.Scan(&name)
			fmt.Print("Enter Initial Balance: ")
			var initialBalance float64
			fmt.Scan(&initialBalance)
			err := createAccount(id, name, initialBalance)
			if err != nil {
				fmt.Printf("Error: %v\n", err)
			} else {
				fmt.Println("Account created successfully.")
			}
		case Deposit:
			fmt.Print("Enter Account ID: ")
			var id int
			fmt.Scan(&id)
			fmt.Print("Enter Amount to Deposit: ")
			var amount float64
			fmt.Scan(&amount)
			err := deposit(id, amount)
			if err != nil {
				fmt.Printf("Error: %v\n", err)
			} else {
				fmt.Println("Deposit successful.")
			}
		case Withdraw:
			fmt.Print("Enter Account ID: ")
			var id int
			fmt.Scan(&id)
			fmt.Print("Enter Amount to Withdraw: ")
			var amount float64
			fmt.Scan(&amount)
			err := withdraw(id, amount)
			if err != nil {
				fmt.Printf("Error: %v\n", err)
			} else {
				fmt.Println("Withdrawal successful.")
			}
		case ViewBalance:
			fmt.Print("Enter Account ID: ")
			var id int
			fmt.Scan(&id)
			balance, err := viewBalance(id)
			if err != nil {
				fmt.Printf("Error: %v\n", err)
			} else {
				fmt.Printf("Account Balance: Rs.%.2f\n", balance)
			}
		case TransactionHistory:
			fmt.Print("Enter Account ID: ")
			var id int
			fmt.Scan(&id)
			history, err := viewTransactionHistory(id)
			if err != nil {
				fmt.Printf("Error: %v\n", err)
			} else {
				fmt.Println("Transaction History:")
				for _, transaction := range history {
					fmt.Println(transaction)
				}
			}
		case Exit:
			fmt.Println("Exiting Bank Transaction System. Goodbye!")
			return
		default:
			fmt.Println("Invalid option. Please choose a valid menu item.")
		}
	}
}

package main

import (
	"errors"
	"fmt"
)

type Account struct {
	Name               string
	ID                 int
	Balance            float64
	TransactionHistory []string
}

const (
	OptionDeposit      = 1
	OptionWithdraw     = 2
	OptionCheckBalance = 3
	OptionViewHistory  = 4
	OptionExit         = 5
)

var accountList []Account

func getAccountByID(id int) (*Account, error) {
	for i := range accountList {
		if accountList[i].ID == id {
			return &accountList[i], nil
		}
	}
	return nil, errors.New("account not found")
}

func deposit(accountID int, amount float64) error {
	if amount <= 0 {
		return errors.New("deposit amount must be greater than zero")
	}

	account, err := getAccountByID(accountID)
	if err != nil {
		return err
	}

	account.Balance += amount
	account.TransactionHistory = append(account.TransactionHistory, fmt.Sprintf("Deposited: $%.2f", amount))
	return nil
}

func withdraw(accountID int, amount float64) error {
	if amount <= 0 {
		return errors.New("withdrawal amount must be greater than zero")
	}

	account, err := getAccountByID(accountID)
	if err != nil {
		return err
	}

	if account.Balance < amount {
		return errors.New("insufficient balance")
	}

	account.Balance -= amount
	account.TransactionHistory = append(account.TransactionHistory, fmt.Sprintf("Withdrew: $%.2f", amount))
	return nil
}

func getBalance(accountID int) (float64, error) {
	account, err := getAccountByID(accountID)
	if err != nil {
		return 0, err
	}
	return account.Balance, nil
}

func getTransactionHistory(accountID int) ([]string, error) {
	account, err := getAccountByID(accountID)
	if err != nil {
		return nil, err
	}
	return account.TransactionHistory, nil
}

func main() {
	accountList = append(accountList, Account{ID: 1, Name: "Alice", Balance: 500.0})
	accountList = append(accountList, Account{ID: 2, Name: "Bob", Balance: 300.0})

	var choice, accountID int
	var amount float64

	for {
		fmt.Println("\n--- Bank Menu ---")
		fmt.Println("1. Deposit Money")
		fmt.Println("2. Withdraw Money")
		fmt.Println("3. Check Balance")
		fmt.Println("4. View Transaction History")
		fmt.Println("5. Exit")
		fmt.Print("Enter your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case OptionDeposit:
			fmt.Print("Enter Account ID: ")
			fmt.Scan(&accountID)
			fmt.Print("Enter Amount to Deposit: ")
			fmt.Scan(&amount)
			if err := deposit(accountID, amount); err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Deposit successful!")
			}
		case OptionWithdraw:
			fmt.Print("Enter Account ID: ")
			fmt.Scan(&accountID)
			fmt.Print("Enter Amount to Withdraw: ")
			fmt.Scan(&amount)
			if err := withdraw(accountID, amount); err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Withdrawal successful!")
			}
		case OptionCheckBalance:
			fmt.Print("Enter Account ID: ")
			fmt.Scan(&accountID)
			balance, err := getBalance(accountID)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Printf("Current Balance: $%.2f\n", balance)
			}
		case OptionViewHistory:
			fmt.Print("Enter Account ID: ")
			fmt.Scan(&accountID)
			history, err := getTransactionHistory(accountID)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Transaction History:")
				for _, entry := range history {
					fmt.Println(entry)
				}
			}
		case OptionExit:
			fmt.Println("Exiting the program. Goodbye!")
			return
		default:
			fmt.Println("Invalid option. Please try again.")
		}
	}
}

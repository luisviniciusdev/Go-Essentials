package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var accountBalance, err = getBalanceFromFile()

	if err != nil {
		fmt.Print(err)
	}

	for {
		printBankMenu()
		userInput := getUserInput()

		switch userInput {
		case 1:
			checkUserBalance(accountBalance)
		case 2:
			depositMoney(&accountBalance)
		case 3:
			withdrawMoney(&accountBalance)
		case 4:
			fmt.Print("\nGood Bye!")
			fmt.Print("\nThanks for choosing Go Bank!")
			return
		default:
			fmt.Print("\nInvalid digit. Please try a valid digit from the list.")
			continue
		}
	}
}

func printBankMenu() {
	fmt.Println("\n\nWelcome to Go Bank!")
	fmt.Println("What do you want to do?")
	fmt.Println("1. Check the balance")
	fmt.Println("2. Deposit money")
	fmt.Println("3. Withdraw money")
	fmt.Println("4. Exit")
}

func getUserInput() int {
	var userInput int
	fmt.Print("\nYour Choice: ")
	fmt.Scan(&userInput)

	return userInput
}

const accountBalanceFile string = "balance.txt"

func writeBalanceToFile(balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile(accountBalanceFile, []byte(balanceText), 0644)
}

func getBalanceFromFile() (float64, error) {
	data, err := os.ReadFile(accountBalanceFile)

	if err != nil {
		return 1518, errors.New("failed to find balance file")
	}

	balanceText := string(data)
	balance, err := strconv.ParseFloat(balanceText, 32)

	if err != nil {
		return 1518, errors.New("failed to parse stored balance value")
	}

	return float64(balance), nil
}

func checkUserBalance(userBalance float64) {
	fmt.Printf("\nYour balance is: %.2f", userBalance)
}

func depositMoney(userBalance *float64) {
	var depositAmount float64
	fmt.Print("\nHow much do you want to deposit? ")
	fmt.Scan(&depositAmount)

	if depositAmount <= 0 {
		print("Invalid Operation. The amount deposited must be greater than 0.")
		return
	}

	*userBalance += depositAmount
	fmt.Printf("\nThe new Balance Amount is: %.2f", *userBalance)
	writeBalanceToFile(*userBalance)
}

func withdrawMoney(userBalance *float64) {
	var withdrawAmount float64
	fmt.Print("\nHow much do you want to withdraw? ")
	fmt.Scan(&withdrawAmount)

	if withdrawAmount > *userBalance {
		print("Invalid Operation. You can`t withdraw more than your current balance.")
		return
	}

	*userBalance -= withdrawAmount
	fmt.Printf("\nThe new Balance Amount is: %.2f", *userBalance)
	writeBalanceToFile(*userBalance)
}

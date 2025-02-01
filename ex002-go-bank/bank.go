package main

import "fmt"

func main() {
	var accountBalance float32 = 1518.00

	printBankMenu()
	userInput := getUserInput()

	if userInput == 1 {
		checkUserBalance(accountBalance)
	}

	if userInput == 2 {
		depositMoney(accountBalance)
	}

	if userInput == 3 {
		withdrawMoney(accountBalance)
	}

	if userInput == 4 || userInput <= 0 || userInput >= 5 {
		fmt.Print("\nGood Bye!")
		fmt.Print("\nThanls for choosing Go Bank!")
		return
	}
}

func printBankMenu() {
	fmt.Println("Welcome to Go Bank!")
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

func checkUserBalance(userBalance float32) {
	fmt.Printf("Your balance is: %.2f", userBalance)
}

func depositMoney(userBalance float32) {
	var depositAmount float32
	fmt.Print("\nHow much do you want to deposit? ")
	fmt.Scan(&depositAmount)

	if depositAmount <= 0 {
		print("Invalid Operation. The amount deposited must be greater than 0.")
		return
	}

	userBalance += depositAmount
	fmt.Printf("\nThe new Balance Amount is: %.2f", userBalance)
}

func withdrawMoney(userBalance float32) {
	var withdrawAmount float32
	fmt.Print("\nHow much do you want to withdraw? ")
	fmt.Scan(&withdrawAmount)

	if withdrawAmount > userBalance {
		print("Invalid Operation. You can`t withdraw more than your current balance.")
		return
	}

	userBalance -= withdrawAmount
	fmt.Printf("\nThe new Balance Amount is: %.2f", userBalance)
}

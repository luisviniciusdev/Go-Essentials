package main

import (
	"fmt"

	"example.com/go-bank/fileops"
	"github.com/Pallinder/go-randomdata"
)

func main() {
	name := randomdata.FirstName(2)
	var accountBalance, err = fileops.GetFloatFromFile(accountBalanceFile)

	if err != nil {
		fmt.Print(err)
	}

	for {

		printBankMenu(name)
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

func getUserInput() int {
	var userInput int
	fmt.Print("\nYour Choice: ")
	fmt.Scan(&userInput)

	return userInput
}

const accountBalanceFile string = "balance.txt"

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
	fileops.WriteFloatToFile(*userBalance, accountBalanceFile)
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
	fileops.WriteFloatToFile(*userBalance, accountBalanceFile)
}

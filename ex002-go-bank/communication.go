package main

import "fmt"

func printBankMenu(userName string) {
	fmt.Printf("\n\nWelcome to Go Bank, %s!\n", userName)
	fmt.Println("What do you want to do?")
	fmt.Println("1. Check the balance")
	fmt.Println("2. Deposit money")
	fmt.Println("3. Withdraw money")
	fmt.Println("4. Exit")
}

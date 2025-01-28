package main

import "fmt"

func main() {
	var revenue, expenses, taxRate float32

	fmt.Print("What is your company`s Revenue? ")
	fmt.Scan(&revenue)

	fmt.Print("What is your company`s Expenses? ")
	fmt.Scan(&expenses)

	fmt.Print("What is your company`s Tax Rate? ")
	fmt.Scan(&taxRate)

	earningsBeforeTax := revenue - expenses
	profit := earningsBeforeTax * (1 - taxRate/100)
	ratio := earningsBeforeTax / profit

	fmt.Printf("\nEarnings Before Tax: $%.2f", earningsBeforeTax)
	fmt.Printf("\nProfit: $%.2f", profit)
	fmt.Printf("\nRatio: %.2f", ratio)
}

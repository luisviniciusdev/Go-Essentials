package main

import "fmt"

func main() {
	revenue := getUserInput("What is your company`s Revenue? ")
	expense := getUserInput("What is your company`s Expenses? ")
	taxRate := getUserInput("What is your company`s Tax Rate? ")

	earningsBeforeTax := calcEarningsBeforeTax(revenue, expense)
	profit := calcProfit(earningsBeforeTax, taxRate)
	ratio := calcRatio(earningsBeforeTax, profit)

	fmt.Printf("\nEarnings Before Tax: $%.2f", earningsBeforeTax)
	fmt.Printf("\nProfit: $%.2f", profit)
	fmt.Printf("\nRatio: %.2f", ratio)
}

func getUserInput(question string) float32 {
	var userInput float32
	fmt.Print(question)
	fmt.Scan(&userInput)
	return userInput
}

func calcEarningsBeforeTax(revenue, expense float32) float32 {
	return revenue - expense
}

func calcProfit(ebt, taxRate float32) float32 {
	return ebt * (1 - taxRate/100)
}

func calcRatio(ebt, profit float32) float32 {
	return ebt / profit
}

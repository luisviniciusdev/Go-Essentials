package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	revenue, err1 := getUserInput("What is your company`s Revenue? ")
	expense, err2 := getUserInput("What is your company`s Expenses? ")
	taxRate, err3 := getUserInput("What is your company`s Tax Rate? ")

	if err1 != nil || err2 != nil || err3 != nil {
		fmt.Print(err1)
		return
	}

	earningsBeforeTax := calcEarningsBeforeTax(revenue, expense)
	profit := calcProfit(earningsBeforeTax, taxRate)
	ratio := calcRatio(earningsBeforeTax, profit)

	fmt.Printf("\nEarnings Before Tax: $%.2f", earningsBeforeTax)
	fmt.Printf("\nProfit: $%.2f", profit)
	fmt.Printf("\nRatio: %.2f", ratio)
	storeResults(earningsBeforeTax, profit, ratio)
}

func getUserInput(question string) (float32, error) {
	var userInput float32
	fmt.Print(question)
	fmt.Scan(&userInput)

	if userInput <= 0 {
		return 0, errors.New("invalid input. the value must be greater than 0")
	}

	return userInput, nil
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

func storeResults(ebt, profit, ratio float32) {
	results := fmt.Sprintf("EarningsBeforeTax: %.2f\nProfit: %.2f\nRatio: %.2f", ebt, profit, ratio)
	os.WriteFile("calculatedResult.txt", []byte(results), 0644)
}

package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	revenue, err := getUserInput("Revenue: ")
	expenses, err := getUserInput("Expenses: ")
	taxRate, err := getUserInput("Tax Rate: ")
	if err != nil {
		panic(err)
		// return
	}

	ebt, profit, ratio := calculateFinancials(revenue,expenses,taxRate)

	fmt.Printf("\nEBT: %.1f\n", ebt)
	fmt.Printf("Profit: %.1f\n", profit)
	fmt.Printf("Profit Ratio: %.3f\n", ratio)
	storeDataToFile(ebt, profit, ratio)
}

func storeDataToFile (ebt, profit, ratio float64) {
	results := fmt.Sprintf("EBT: %.1f\nPROFIT: %.1f\nRATIO: %.3f\n", ebt, profit, ratio)
	os.WriteFile("results.txt", []byte(results), 0644)
}

func getUserInput (str string) (float64, error) {
	// fmt.Println(str)
	var r float64
	fmt.Print(str)
	fmt.Scan(&r)

	if r <= 0 {
		return 0, errors.New("Number must be positive")
	}

	return r, nil
}

func calculateFinancials (r, e, tr float64) (ebt, profit, ratio float64) {
	ebt = r - e
	profit = ebt * (1-tr/100)
	ratio = ebt / profit
	return
}
package main

import "fmt"

func main() {
	revenue := getUserInput("Revenue: ")
	expenses := getUserInput("Expenses: ")
	taxRate := getUserInput("Tax Rate: ")

	ebt, profit, profitRatio := calculateFinancials(revenue,expenses,taxRate)

	fmt.Printf("\nEBT: %.1f\n", ebt)
	fmt.Printf("Profit: %.1f\n", profit)
	fmt.Printf("Profit Ratio: %.3f\n", profitRatio)
}

func getUserInput (str string) (r float64) {
	fmt.Print(str)
	fmt.Scan(&r)
	return
}

func calculateFinancials (r, e, tr float64) (ebt, profit, profitRatio float64) {
	ebt = r - e
	profit = ebt * (1-tr/100)
	profitRatio = ebt / profit
	return
}
package main

import "fmt"

func main() {
	var revenue, expenses, taxRate float64
	
	fmt.Print("Revenue: ")
	fmt.Scan(&revenue)
	fmt.Print("Expenses: ")
	fmt.Scan(&expenses)
	fmt.Print("Tax Rate: ")
	fmt.Scan(&taxRate)

	ebt := revenue - expenses
	profit := ebt * (1-taxRate/100)
	profitRatio := ebt / profit

	fmt.Printf("\nEBT: %.2f\n", ebt)
	fmt.Printf("Profit: %.2f\n", profit)
	fmt.Printf("Profit Ratio: %.2f\n", profitRatio)
}
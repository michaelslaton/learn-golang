package main

import "fmt"

func main(){
	var productNames [4]string = [4]string{"A Book"}
	prices := [4]float64{10.99, 9.99, 45.99, 20.0}
	fmt.Println(prices)

	productNames[2] = "A Carpet"

	fmt.Println(productNames)

	fmt.Println(prices[2])

	// featuredPrices := prices[1:3]
	// fmt.Println(featuredPrices)

	// featuredPrices2 := prices[:3]
	// fmt.Println(featuredPrices2)

	// featuredPrices3 := prices[1:]
	// fmt.Println(featuredPrices3)

	// featuredPrices3 := prices[1:4]
	// fmt.Println(featuredPrices3)

	featuredPrices := prices[1:]
	highlightedPrices := featuredPrices[:1]
	fmt.Println(highlightedPrices)
}
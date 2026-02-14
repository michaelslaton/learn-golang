package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 6.5
	var investmentAmount, years float64
	expectedReturnRate := 5.5

	fmt.Print("Amount to invest: ")
	fmt.Scan(&investmentAmount)
	fmt.Print("Years to invest: ")
	fmt.Scan(&years)
	fmt.Print("Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	// aThing := make([]float64,5)
	// aThing = append(aThing, 65.0)
	// aThing[3] = 4
	
	// for i, val := range aThing {
	// 	if val == 65 {
	// 		fmt.Printf("Found it at %d\n", i)
	// 	} else {
	// 		fmt.Println("No dice")
	// 	}
	// }

	var futureValue = investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	
	formattedFV := fmt.Sprintf("\nFuture Value: %.2f\n", futureValue)
	formattedRFV := fmt.Sprintf("Future Real Value: %.2f\n", futureRealValue)
	
	outputText(formattedFV + formattedRFV)
	// fmt.Printf(
	// `Future Value: %.2f
	// Future Real Value: %.2f`,
	// futureValue, futureRealValue)
	// fmt.Print(formattedFV,formattedRFV)
}

func outputText(text string){
	fmt.Print(text)
}
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

	fmt.Printf("Future Value: %v\nFuture Real Value: %v\n", futureValue, futureRealValue)
}
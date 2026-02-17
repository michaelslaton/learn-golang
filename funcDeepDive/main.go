package main

import "fmt"

func main() {
	// numbers := []int{1,10,15}
	sum := sumup(1,10,15, -5)

	fmt.Println(sum)

}

func sumup(startingValue int, numbers ...int) int {
	sum := 0

	for _, v := range numbers {
		sum += v
	}

	return sum
}
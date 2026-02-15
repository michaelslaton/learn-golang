package main

import "fmt"

func main () {

	mySlice := make([] int, 5)

	for i := range mySlice {
		mySlice[i] = i
	}

	for i:=2; i != len(mySlice); i++ {
		print(mySlice[i])
	}

	fmt.Print(mySlice)
}
package main

import "fmt"

func main () {

	mySlice := make([] int, 5)

	for i,v := range mySlice {
		mySlice[i] = i
		fmt.Println(v)
	}

	fmt.Print(mySlice)
}
package main

import "fmt"

func main () {
	age := 32

	var agePointer *int 
	agePointer = &age

	fmt.Println("Age:", *agePointer)

	convertToAdultYears(agePointer)
	fmt.Println(age)
}

func convertToAdultYears(age *int) {
	// return *age - 18
	*age = *age - 18
}
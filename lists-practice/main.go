package main

import "fmt"

func main() {
	fmt.Println("1---------------------------------------------------")
	hobbies := [3]string{"photography", "gaming", "collecting"}
	fmt.Println(hobbies)

	fmt.Println("2---------------------------------------------------")
	fmt.Println(hobbies[0])
	fmt.Println(hobbies[1:])

	fmt.Println("3---------------------------------------------------")
	firstSecondHobbies := hobbies[:2]
	// firstSecondHobbies2 := hobbies[0:2]
	fmt.Println(firstSecondHobbies)

	fmt.Println("4---------------------------------------------------")
	firstSecondHobbies = firstSecondHobbies[1:3]
	fmt.Println(firstSecondHobbies)

	fmt.Println("5---------------------------------------------------")
	courseGoals := []string{"Learn Go", "Get Fulltime at Uber"}
	fmt.Println(courseGoals)

	fmt.Println("6---------------------------------------------------")
	courseGoals[1] = "To make this Eng Boost course easier"
	courseGoals = append(courseGoals, "...still to get on full time at Uber")
	fmt.Println(courseGoals)

	fmt.Println("7---------------------------------------------------")

	type product struct {
		id int
		title string
		price float64
	}

	products := []product{
		{
			0,
			"Egg",
			5.99,
		},
		{
			1,
			"Milk",
			4.99,
		}}
	fmt.Println(products)
	newProduct := product{
		2,
		"Cereal",
		1.00,
	}
	products = append(products, newProduct)
	fmt.Println(products)
}

// Time to practice what you learned!

// 1) Create a new array (!) that contains three hobbies you have
// 		Output (print) that array in the command line.
// 2) Also output more data about that array:
//		- The first element (standalone)
//		- The second and third element combined as a new list
// 3) Create a slice based on the first element that contains
//		the first and second elements.
//		Create that slice in two different ways (i.e. create two slices in the end)
// 4) Re-slice the slice from (3) and change it to contain the second
//		and last element of the original array.
// 5) Create a "dynamic array" that contains your course goals (at least 2 goals)
// 6) Set the second goal to a different one AND then add a third goal to that existing dynamic array
// 7) Bonus: Create a "Product" struct with title, id, price and create a
//		dynamic list of products (at least 2 products).
//		Then add a third product to the existing list of products.
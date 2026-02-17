package main

import "fmt"

// import "fmt"

func main(){
	userNames := make([]string, 2, 5)

	userNames[0] = "Julie"
	userNames[1] = "Barbara"

	userNames = append(userNames, "Max")
	userNames = append(userNames, "Manuel")

	fmt.Println(userNames)
}
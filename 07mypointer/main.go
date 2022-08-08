package main

import "fmt"

func main() {
	fmt.Println("welcome to pointer study")

	// var ptr *int
	// fmt.Println("Value of pointer is ", ptr)

	myNumber := 23

	var ptr = &myNumber
	fmt.Println("vlaue of pointer is ", ptr)
	fmt.Println("vlaue of pointer is ", *ptr)
}

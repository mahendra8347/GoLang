package main

import "fmt"

func main() {
	fmt.Println("If else in golang")

	logincount := 23
	var result string

	if logincount < 10 {
		result = "regulat user"
	} else if logincount > 10 {
		result = "non regulate user"
	} else {
		result = "non regulate user"
	}

	fmt.Println(result)

	if 9%2 == 0 {
		fmt.Println("Number is even")
	} else {
		fmt.Println("Number is odd")
	}

	if num := 3; num < 10 {
		fmt.Println("Number is less then 10")
	} else {
		fmt.Println("Number is greater then 10")
	}

}

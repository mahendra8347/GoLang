package main

import "fmt"

func main() {
	fmt.Println("function tutorial")
	greeter()

	result := adder(3, 5)
	fmt.Println("Result is ", result)

	proint, prostring := proAdder(2, 5, 6, 8)
	fmt.Printf("Result is %v and string is %v", proint, prostring)

}

func adder(valOne int, vaTwo int) int {
	return valOne + vaTwo
}

func proAdder(vals ...int) (int, string) {
	sum := 0
	for _, val := range vals {
		sum += val
	}
	return sum, "Hi sum is"
}

func greeter() {
	fmt.Println("Namaste from golang")
}

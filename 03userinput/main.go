package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "welcome to user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter a value")

	// comma ok // err ok
	input, _ := reader.ReadString('\n')
	fmt.Println("You have entered: ", input)

}

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Switch and case in golong")

	rand.Seed(time.Now().Unix())
	diceNumber := rand.Intn(6) + 1
	fmt.Println("Value of dice is ", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("dice value is one and you can open")
	case 2:
		fmt.Println("You can move 2 sport")
	case 3:
		fmt.Println("You can move 3 sport")
		fallthrough
	case 4:
		fmt.Println("You can move 4 sport")
		fallthrough
	case 5:
		fmt.Println("You can move 5 sport")
	case 6:
		fmt.Println("You can move 6 sport and roll dice again")
	}

}

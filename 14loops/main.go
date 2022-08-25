package main

import "fmt"

func main() {
	fmt.Println("loops in golong")

	days := []string{"Sunday", "Tuesday", "Wednesday", "Friday", "Saturday"}

	fmt.Println(days)

	for d := 0; d < len(days); d++ {
		fmt.Println(days[d])
	}

	for i := range days {
		fmt.Println(days[i])
	}

	for index, day := range days {
		fmt.Printf("Index is %v and value is %v\n ", index, day)
	}

	rougueValue := 1

	for rougueValue < 10 {
		if rougueValue == 2 {
			goto lco
		}
		if rougueValue == 5 {
			break
		}
		if rougueValue == 3 {
			rougueValue++
			continue
		}
		fmt.Println("Value is ", rougueValue)
		rougueValue++
	}

lco:
	fmt.Println("Jumped to loc")
}

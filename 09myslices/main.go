package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("welcome to slive tutorial")

	var fruitList = []string{"Apple", "Tometo", "Peach"}
	fmt.Printf("Type of fruit list is %T \n", fruitList)

	fruitList = append(fruitList, "Mango", "Banana")
	fmt.Println(fruitList)

	fruitList = append(fruitList[1:3])
	fmt.Println(fruitList)

	highScores := make([]int, 4)

	highScores[0] = 235
	highScores[1] = 569
	highScores[2] = 369
	highScores[3] = 869

	highScores = append(highScores, 555, 693)
	fmt.Println(highScores)

	sort.Ints(highScores)
	fmt.Println(highScores)
	fmt.Println(sort.IntsAreSorted(highScores))

	var courses = []string{"reactjs", "javascript", "python", "swift", "ruby"}
	fmt.Println(courses)
	index := 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)
}

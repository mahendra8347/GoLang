package main

import "fmt"

func main() {
	fmt.Println("welcome to maps study")

	languages := make(map[string]string)

	languages["JS"] = "Javascript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"
	fmt.Println("List of languages", languages)
	fmt.Println("JS shorts for : ", languages["JS"])

	delete(languages, "RB")
	fmt.Println("List of languages", languages)

	// loop are interresting in golang

	for key, value := range languages {
		fmt.Printf("for key %v, value is %v \n", key, value)
	}
}

package main

import "fmt"

func main() {
	fmt.Println("Welcomme to struct study")

	// no inheritance in golang; no super or parent

	mahnedra := User{"Mahendra", "mlparmar@gmail.com", true, 23}
	fmt.Println(mahnedra)
	fmt.Printf("Mahendra details are: %+v \n", mahnedra)
	fmt.Printf("Name is: %v", mahnedra.Name)
	mahnedra.GetStatus()
	mahnedra.NewMail()
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("\nIs user active: ", u.Status)
}

func (u User) NewMail() {
	u.Email = "test@go.dev"
	fmt.Println("Email of this user is ", u.Email)
}

package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("welcome to time study of golang")
	presentTime := time.Now()
	fmt.Println(presentTime)

	fmt.Println((presentTime.Format("01-02-2006 15:04:05 Monday")))

	createdDate := time.Date(2022, time.August, 06, 10, 52, 0, 0, time.UTC)

	fmt.Println(createdDate)
	fmt.Println(createdDate.Format("01-02-2006 15:04:05 Monday"))

}

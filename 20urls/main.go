package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=abc459"

func main() {
	fmt.Println("welcome to hendling URLs in golang")
	fmt.Println(myurl)

	//parsing

	result, _ := url.Parse(myurl)

	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.RawQuery)
	fmt.Println(result.Port())

	qparams := result.Query()
	fmt.Printf("data of query is %T\n", qparams)
	fmt.Println(qparams)
	fmt.Println(qparams["coursename"])

	for _, val := range qparams {
		fmt.Println("param is: ", val)
	}

	partsOfURL := &url.URL{
		Scheme:  "http",
		Host:    "loc.dev",
		RawPath: "user=hitesh",
		Path:    "/tutcss",
	}

	anotherUrl := partsOfURL.String()
	fmt.Println(anotherUrl)
}

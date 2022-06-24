package main

import (
	"fmt"
	"log"
	"net/url"
)

func main() {
	URL, err := url.Parse("https://google.com/search/?query=asdjk")

	if err != nil {
		log.Fatal("Error occured")
	}

	fmt.Println("Path: ", URL.Path)
	fmt.Println("Host: ", URL.Host)
	fmt.Println("Scheme: ", URL.Scheme)
	fmt.Println("Query: ", URL.RawQuery)
}

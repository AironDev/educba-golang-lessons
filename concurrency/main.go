package main

import (
	"fmt"
	"time"
)

func shout(text string) {
	for i := 0; i < 5; i++ {
		fmt.Println(text)
		time.Sleep(1 * time.Second)
	}
}
func main() {
	go shout("Eyaaah")
	shout("Ewoooo")
}

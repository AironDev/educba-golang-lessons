package main

import "fmt"

func main() {
	var i int = 10
LOOP:
	for i < 30 {
		if i%3 == 0 {
			i += 2
			goto LOOP
		}
		fmt.Printf("Numbers not divisible by 3 is %d\n", i)
		i++
	}
}

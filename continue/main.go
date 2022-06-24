package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	scores := []int{40, 50, 13, 90, 80, 60, 12, 39, 21, 60}
	fmt.Println("Old scores: ", scores)
MAKEUP:
	for index, score := range scores {
		if score > 40 {
			continue
		}

		rand.Seed(time.Now().UnixNano())
		scores[index] += rand.Intn(10)

		goto MAKEUP
	}

	fmt.Println("New scores: ", scores)
}

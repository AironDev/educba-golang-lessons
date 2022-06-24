package loops

import "fmt"

func SumLoop() int {
	result := 0
	for i := 0; i < 10; i++ {
		result += i
	}
	return result

}

func SumArray(arr []int) int {
	var result int
	for _, num := range arr {
		result += num
	}
	return result
}

func ListElements(elements []string) {
	for index, value := range elements {
		fmt.Println(index, value)
	}

}

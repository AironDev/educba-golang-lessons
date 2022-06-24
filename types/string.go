package types

import "fmt"

type Names []string

func (n Names) Print() {
	fmt.Println(n)
}

func (n Names) AddName(name string) Names {
	return append(n, name)
}

func RollCall(names Names) {
	for _, name := range names {
		fmt.Println(name)
	}
}

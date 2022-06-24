package main

import (
	"fmt"
	"github.com/airondev/educba/loops"
	"github.com/airondev/educba/strings"
	"github.com/airondev/educba/types"
)

func main() {
	// sum plain loop
	//fmt.Println(loops.SumLoop())

	//sum array values
	//values := []int{1, 2, 3, 8, 10}
	//fmt.Println(loops.SumArray(values))

	// string methods
	fmt.Println("Length is ", strings.StringLen("Aaron"))
	fmt.Println("Uppercase", strings.StringToUpper("Aaron"))
	fmt.Println("Split", strings.StringSplit("Aaron, John, Sam, doe", ","))

	strings.StringReplace()

	var names = []string{"John", "Doe"}
	fmt.Println("Full name is", strings.StringJoin(names))

	// print array values with index
	loops.ListElements([]string{"Helium", "Beryllium", "Hydrogen"})

	myNames := types.Names{"John", "Ken", "Sam", "Edet", "Mimmie"}
	myNames = myNames.AddName("Kennedy")
	myNames.Print()
}

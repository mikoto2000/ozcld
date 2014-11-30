package main

import (
	cld "../../ozcld"
	"fmt"
)

func main() {
	methods := cld.CreateMethodsFromStrings([]string{})

	methods.Add(cld.CreateMethodFromString("Testing1"))
	methods.Add(cld.CreateMethodFromString("Testing2"))
	methods.Add(cld.CreateMethodFromString("Testing3"))

	fmt.Println(methods.ToDot())
}

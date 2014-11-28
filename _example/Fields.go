package main

import (
	cld "../../ozcld"
	"fmt"
)

func main() {
	fields := cld.CreateFieldsFromStrings([]string{})

	fields.Add(cld.CreateFieldFromString("Testing1"))
	fields.Add(cld.CreateFieldFromString("Testing2"))
	fields.Add(cld.CreateFieldFromString("Testing3"))

	fmt.Println(fields.ToDot())
}

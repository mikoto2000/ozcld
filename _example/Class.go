package main

import (
	cld "../../ozcld"
	"fmt"
)

func main() {
	class := cld.CreateClassFromDefs("", "TestClass", nil, nil)

	class.AddFieldFromString("- FieldTesting1 : string")
	class.AddFieldFromString("- FieldTesting2 : string")
	class.AddFieldFromString("- FieldTesting3 : string")

	class.AddMethodFromString("+ MethodTesting1() : string")
	class.AddMethodFromString("+ MethodTesting2() : string")
	class.AddMethodFromString("+ MethodTesting3() : string")

	fmt.Println(class.ToDot())
}

package main

import (
	cld "../../ozcld"
	"fmt"
)

func main() {

	class1 := cld.CreateClassFromDefs("", "TestClassInNamespace1", nil, nil)

	class1.AddFieldFromString("- FieldTesting1 : string")
	class1.AddFieldFromString("- FieldTesting2 : string")
	class1.AddFieldFromString("- FieldTesting3 : string")

	class1.AddMethodFromString("+ MethodTesting1() : string")
	class1.AddMethodFromString("+ MethodTesting2() : string")
	class1.AddMethodFromString("+ MethodTesting3() : string")

	class2 := cld.CreateClassFromDefs("", "TestClassInNamespace2", []string{"aaa"}, []string{"bbb"})

	classes := make([]*cld.Class, 2)
	classes[0] = class1
	classes[1] = class2

	namespace := cld.CreateNamespace("TestNamespace", classes)

	fmt.Println(namespace.ToDot())
}

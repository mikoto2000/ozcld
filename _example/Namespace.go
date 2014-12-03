package main

import (
	cld "../../ozcld"
	"fmt"
)

func main() {

	class1 := cld.CreateClassFromDefs("", "TestClassInInnerNamespace", nil, nil)

	class1.AddFieldFromString("- FieldTesting1 : string")
	class1.AddFieldFromString("- FieldTesting2 : string")
	class1.AddFieldFromString("- FieldTesting3 : string")

	class1.AddMethodFromString("+ MethodTesting1() : string")
	class1.AddMethodFromString("+ MethodTesting2() : string")
	class1.AddMethodFromString("+ MethodTesting3() : string")

	class2 := cld.CreateClassFromDefs("", "TestClassInOuterNamespace", []string{"aaa"}, []string{"bbb"})

	innerNamespace := cld.CreateNamespace("TestInnerNamespace", nil, nil)
	innerNamespace.AddClass(class1)
	outerNamespace := cld.CreateNamespace("TestOuterNamespace", nil, []*cld.Namespace{innerNamespace})
	outerNamespace.AddClass(class2)

	fmt.Println(outerNamespace.ToDot())
}

package ozcld

import (
	cld "../../ozcld"
	"testing"
)

func TestMain(t *testing.T) {
	methodDefs := []string{"- method01() : string",
		"- method02() : string"}

	//  インタフェース(ステレオタイプあり、フィールドなし)
	myIf := cld.CreateClassFromDefs("Interface", "TestInterface", nil, methodDefs)

	actual1 := myIf.ToDot()
	expected1 := "main_TestInterface [label = \"{\\<\\<Interface\\>\\>\\nTestInterface||- method01() : string\\l- method02() : string\\l}\"];"

	if expected1 != actual1 {
		t.Errorf("got\n \"%s\"\nbut want\n \"%s\"", actual1, expected1)
	}

	// クラス(ステレオタイプなし、フィールドあり)
	fieldDefs := []string{"- field01 : string",
		"- field02 : string"}

	myClass := cld.CreateClassFromDefs("", "TestClass", fieldDefs, methodDefs)

	actual2 := myClass.ToDot()
	expected2 := "main_TestClass [label = \"{TestClass|- field01 : string\\l- field02 : string\\l|- method01() : string\\l- method02() : string\\l}\"];"

	if expected2 != actual2 {
		t.Errorf("got\n \"%s\"\nbut want\n \"%s\"", actual2, expected2)
	}
}

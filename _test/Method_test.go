package ozcld

import (
	cld "../../ozcld"
	"testing"
)

func TestMethod(t *testing.T) {
	def := "- method01() : string"
	method := cld.CreateMethodFromString(def)

	actual := method.ToDot()
	expected := def + "\\l"

	if expected != actual {
		t.Errorf("got \"%s\" but want \"%s\"", actual, expected)
	}
}

func TestMethods(t *testing.T) {
	defs := []string{"- method01() : string",
		"- method02() : string"}
	methods := cld.CreateMethodsFromStrings(defs)

	actual := methods.ToDot()
	expected := defs[0] + "\\l" + defs[1] + "\\l"

	if expected != actual {
		t.Errorf("got \"%s\" but want \"%s\"", actual, expected)
	}
}

func TestMethodsAdd(t *testing.T) {
	defs := []string{"- method01() : string"}
	def := "- method02() : string"
	methods := cld.CreateMethodsFromStrings(defs)
	methods.Add(cld.CreateMethodFromString(def))

	actual := methods.ToDot()
	expected := defs[0] + "\\l" + def + "\\l"

	if expected != actual {
		t.Errorf("got \"%s\" but want \"%s\"", actual, expected)
	}
}

func getMethodsDefs() []string {
	return []string{"- m1() : string", "- m2() : string", "- m3() : string"}
}

func getMethods() *cld.Methods {
	return cld.CreateMethodsFromStrings(getMethodsDefs())
}

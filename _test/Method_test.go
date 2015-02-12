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

func getMethodsDefs() []string {
	return []string{"- m1() : string", "- m2() : string", "- m3() : string"}
}

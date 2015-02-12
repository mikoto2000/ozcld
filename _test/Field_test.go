package ozcld

import (
	cld "../../ozcld"
	"testing"
)

func TestField(t *testing.T) {
	def := "- field : string"
	field := cld.CreateFieldFromString(def)

	actual := field.ToDot()
	expected := def + "\\l"

	if expected != actual {
		t.Errorf("got \"%s\" but want \"%s\"", actual, expected)
	}
}

func getFieldsDefs() []string {
	return []string{"- f1 : string", "- f2 : string", "- f3 : string"}
}

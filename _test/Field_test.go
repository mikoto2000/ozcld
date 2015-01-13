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

func TestFields(t *testing.T) {
	defs := []string{"- field01 : string",
		"- field02 : string"}
	fields := cld.CreateFieldsFromStrings(defs)

	actual := fields.ToDot()
	expected := defs[0] + "\\l" + defs[1] + "\\l"

	if expected != actual {
		t.Errorf("got \"%s\" but want \"%s\"", actual, expected)
	}
}

func TestFieldsAdd(t *testing.T) {
	defs := []string{"- field01 : string"}
	def := "- field02 : string"
	fields := cld.CreateFieldsFromStrings(defs)
	fields.Add(cld.CreateFieldFromString(def))

	actual := fields.ToDot()
	expected := defs[0] + "\\l" + def + "\\l"

	if expected != actual {
		t.Errorf("got \"%s\" but want \"%s\"", actual, expected)
	}
}

func getFieldsDefs() []string {
	return []string{"- f1 : string", "- f2 : string", "- f3 : string"}
}

func getFields() *cld.Fields {
	return cld.CreateFieldsFromStrings(getFieldsDefs())
}

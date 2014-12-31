package ozcld

import (
	cld "../../ozcld"
	"testing"
)

func TestRelation(t *testing.T) {
	relation := cld.CreateRelation("RelationName", cld.RELATION_INHERIT, "TestClass1", "TestClass2", "fromMultiplicity", "toMultiplicity")

	actual1 := relation.ToDot()
	expected1 := "edge [style = \"solid\", arrowhead = \"onormal\"];\nTestClass1 -> TestClass2[label = \"RelationName\",taillabel = \"fromMultiplicity\",headlabel = \"toMultiplicity\"];"

	if expected1 != actual1 {
		t.Errorf("got\n \"%s\"\nbut want\n \"%s\"", actual1, expected1)
	}
}

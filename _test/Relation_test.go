package ozcld

import (
	cld "../../ozcld"
	"testing"
)

func TestRelation(t *testing.T) {
	c1 := "TestClass1"
	c2 := "TestClass2"
	relation := cld.CreateRelation("RelationName", cld.RELATION_INHERIT, c1, c2, "fromMultiplicity", "toMultiplicity")

	actual1 := relation.ToDot()
	expected1 := "edge [style = \"solid\", arrowhead = \"onormal\"];\nTestClass1 -> TestClass2[label = \"RelationName\",taillabel = \"fromMultiplicity\",headlabel = \"toMultiplicity\"];"

	if expected1 != actual1 {
		t.Errorf("got\n \"%s\"\nbut want\n \"%s\"", actual1, expected1)
	}
}

package ozcld

import (
	cld "../../ozcld"
	"testing"
)

func TestRelation(t *testing.T) {
	c1 := cld.CreateClassFromDefs("", "TestClass1", nil, nil)
	c2 := cld.CreateClassFromDefs("", "TestClass2", nil, nil)
	relation := cld.CreateRelation("RelationName", cld.RELATION_INHERIT, c1, c2, "fromMultiplicity", "toMultiplicity")

	actual1 := relation.ToDot()
	expected1 := "edge [style = \"solid\", arrowhead = \"onormal\"];\nmain_TestClass1 -> main_TestClass2[label = \"RelationName\",taillabel = \"fromMultiplicity\",headlabel = \"toMultiplicity\"];"

	if expected1 != actual1 {
		t.Errorf("got\n \"%s\"\nbut want\n \"%s\"", actual1, expected1)
	}
}

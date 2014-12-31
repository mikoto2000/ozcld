package ozcld

import (
	cld "../../ozcld"
	"testing"
)

func TestNamespace(t *testing.T) {
	namespace := cld.CreateNamespace("TestNamespace", []*cld.Class{
		cld.CreateClassFromDefs("", "TestClass1", getFieldsDefs(), getMethodsDefs()),
		cld.CreateClassFromDefs("", "TestClass2", getFieldsDefs(), getMethodsDefs())}, nil)

	actual1 := namespace.ToDot()
	expected1 := "subgraph cluster_TestNamespace {\nlabel = \"TestNamespace\";\nTestClass1 [label = \"{TestClass1|- f1 : string\\l- f2 : string\\l- f3 : string\\l|- m1() : string\\l- m2() : string\\l- m3() : string\\l}\"];\nTestClass2 [label = \"{TestClass2|- f1 : string\\l- f2 : string\\l- f3 : string\\l|- m1() : string\\l- m2() : string\\l- m3() : string\\l}\"];\n}"

	if expected1 != actual1 {
		t.Errorf("got\n \"%s\"\nbut want\n \"%s\"", actual1, expected1)
	}
}

func TestNamespaceAddClass(t *testing.T) {
	namespace := cld.CreateNamespace("TestNamespace", nil, nil)
	namespace.AddClass(cld.CreateClassFromDefs("", "TestClass1", getFieldsDefs(), getMethodsDefs()))
	namespace.AddClass(cld.CreateClassFromDefs("", "TestClass2", getFieldsDefs(), getMethodsDefs()))

	actual1 := namespace.ToDot()
	expected1 := "subgraph cluster_TestNamespace {\nlabel = \"TestNamespace\";\nTestClass1 [label = \"{TestClass1|- f1 : string\\l- f2 : string\\l- f3 : string\\l|- m1() : string\\l- m2() : string\\l- m3() : string\\l}\"];\nTestClass2 [label = \"{TestClass2|- f1 : string\\l- f2 : string\\l- f3 : string\\l|- m1() : string\\l- m2() : string\\l- m3() : string\\l}\"];\n}"

	if expected1 != actual1 {
		t.Errorf("got\n \"%s\"\nbut want\n \"%s\"", actual1, expected1)
	}
}

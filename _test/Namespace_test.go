package ozcld

import (
	cld "../../ozcld"
	"testing"
)

func TestNamespace(t *testing.T) {
	namespace1 := cld.CreateNamespace("TestNamespace", []*cld.Class{
		cld.CreateClassFromDefs("", "TestClass1", getFieldsDefs(), getMethodsDefs()),
		cld.CreateClassFromDefs("", "TestClass2", getFieldsDefs(), getMethodsDefs())}, nil)

	actual1 := namespace1.ToDot()
	expected1 := "subgraph cluster_TestNamespace {\nlabel = \"TestNamespace\";\nTestClass1 [label = \"{TestClass1|- f1 : string\\l- f2 : string\\l- f3 : string\\l|- m1() : string\\l- m2() : string\\l- m3() : string\\l}\"];\nTestClass2 [label = \"{TestClass2|- f1 : string\\l- f2 : string\\l- f3 : string\\l|- m1() : string\\l- m2() : string\\l- m3() : string\\l}\"];\n}"

	if expected1 != actual1 {
		t.Errorf("got\n \"%s\"\nbut want\n \"%s\"", actual1, expected1)
	}

	namespace2 := cld.CreateNamespace("TestNamespace", []*cld.Class{
		cld.CreateClassFromDefs("", "TestClass1", getFieldsDefs(), getMethodsDefs()),
		cld.CreateClassFromDefs("", "TestClass2", getFieldsDefs(), getMethodsDefs())},
		[]*cld.Namespace{getTestNamespace("TestNamespace_innerNamespace")})

	actual2 := namespace2.ToDot()
	expected2 := "subgraph cluster_TestNamespace {\nlabel = \"TestNamespace\";\nsubgraph cluster_TestNamespace_innerNamespace {\nlabel = \"TestNamespace_innerNamespace\";\nTestClass1 [label = \"{TestClass1|- f1 : string\\l- f2 : string\\l- f3 : string\\l|- m1() : string\\l- m2() : string\\l- m3() : string\\l}\"];\nTestClass2 [label = \"{TestClass2|- f1 : string\\l- f2 : string\\l- f3 : string\\l|- m1() : string\\l- m2() : string\\l- m3() : string\\l}\"];\n}\nTestClass1 [label = \"{TestClass1|- f1 : string\\l- f2 : string\\l- f3 : string\\l|- m1() : string\\l- m2() : string\\l- m3() : string\\l}\"];\nTestClass2 [label = \"{TestClass2|- f1 : string\\l- f2 : string\\l- f3 : string\\l|- m1() : string\\l- m2() : string\\l- m3() : string\\l}\"];\n}"

	if expected2 != actual2 {
		t.Errorf("got\n \"%s\"\nbut want\n \"%s\"", actual2, expected2)
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

func getTestNamespace(name string) *cld.Namespace {
	namespace := cld.CreateNamespace(name, []*cld.Class{
		cld.CreateClassFromDefs("", "TestClass1", getFieldsDefs(), getMethodsDefs()),
		cld.CreateClassFromDefs("", "TestClass2", getFieldsDefs(), getMethodsDefs())}, nil)
	return namespace
}

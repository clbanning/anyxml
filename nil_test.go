package anyxml

import (
	"fmt"
	"testing"
)

func TestNilHeader2(t *testing.T) {
	fmt.Println("\n----------------  nil_test.go ...")
}

func TestNilMap(t *testing.T) {
	checkval := "<root/>"
	xmlout, err := Xml(nil, "root")
	if err != nil {
		t.Fatal(err)
	}
	if string(xmlout) != checkval {
		fmt.Println(string(xmlout), "!=", checkval)
		t.Fatal()
	}

	checkval = "   <root/>"
	xmlout, err = XmlIndent(nil, "   ", "  ", "root")
	if err != nil {
		t.Fatal(err)
	}
	if string(xmlout) != checkval {
		fmt.Println(string(xmlout), "!=", checkval)
		t.Fatal()
	}

	// use Go XML marshal syntax for empty element"
	UseGoEmptyElementSyntax = true

	checkval = "<root></root>"
	xmlout, err = Xml(nil, "root")
	if err != nil {
		t.Fatal(err)
	}
	if string(xmlout) != checkval {
		fmt.Println(string(xmlout), "!=", checkval)
		t.Fatal()
	}

	checkval = "   <root></root>"
	xmlout, err = XmlIndent(nil, "   ", "  ", "root")
	if err != nil {
		t.Fatal(err)
	}
	if string(xmlout) != checkval {
		fmt.Println(string(xmlout), "!=", checkval)
		t.Fatal()
	}
}

func TestNilValue(t *testing.T) {
	val := map[string]interface{}{"toplevel": nil}
	checkval := "<root><toplevel/></root>"

	UseGoEmptyElementSyntax = false
	xmlout, err := Xml(val, "root")
	if err != nil {
		t.Fatal(err)
	}
	if string(xmlout) != checkval {
		fmt.Println(string(xmlout), "!=", checkval)
		t.Fatal()
	}

	checkval = `   <root>
     <toplevel/>
   </root>`
	xmlout, err = XmlIndent(val, "   ", "  ", "root")
	if err != nil {
		t.Fatal(err)
	}
	if string(xmlout) != checkval {
		fmt.Println(string(xmlout), "!=", checkval)
		t.Fatal()
	}

	UseGoEmptyElementSyntax = true
	checkval = "<root><toplevel></toplevel></root>"
	xmlout, err = Xml(val, "root")
	if err != nil {
		t.Fatal(err)
	}
	if string(xmlout) != checkval {
		fmt.Println(string(xmlout), "!=", checkval)
		t.Fatal()
	}

	checkval = `   <root>
     <toplevel></toplevel>
   </root>`
	xmlout, err = XmlIndent(val, "   ", "  ", "root")
	if err != nil {
		t.Fatal(err)
	}
	if string(xmlout) != checkval {
		fmt.Println(string(xmlout), "!=", checkval)
		t.Fatal()
	}
}

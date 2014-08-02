package anyxml

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestAnyXmlHeader(t *testing.T) {
	fmt.Println("\n----------------  anyxml_test.go ...\n")
}

var anydata = []byte(`[
    {
        "somekey": "somevalue"
    },
    {
        "somekey": "somevalue"
    },
    {
        "somekey": "somevalue",
        "someotherkey": "someothervalue"
    },
	"a string",
	3.14159625,
	true
]`)

type MyStruct struct {
	Somekey string  `xml:"somekey"`
	B       float32 `xml:"floatval"`
}

func TestXml(t *testing.T) {
	var i interface{}
	err := json.Unmarshal(anydata, &i)
	x, err := Xml(i)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("[]->x:", string(x))

	a := []interface{}{ "try", "this", 3.14159265, true }
	x, err = Xml(a)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("a->x:", string(x))

	x, err = Xml(3.14159625)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("f->x:", string(x))

	s := MyStruct{"somevalue", 3.14159625}
	x, err = Xml(s)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("s->x:", string(x))
}

func TestXmlIndent(t *testing.T) {
	var i interface{}
	err := json.Unmarshal(anydata, &i)
	x, err := XmlIndent(i, "", "  ")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("[]->x:\n", string(x))

	a := []interface{}{ "try", "this", 3.14159265, true }
	x, err = XmlIndent(a, "", "  ")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("a->x:\n", string(x))

	x, err = XmlIndent(3.14159625, "", "  ")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("f->x:\n", string(x))

	s := MyStruct{"somevalue", 3.14159625}
	x, err = XmlIndent(s, "", "  ")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("s->x:\n", string(x))
}

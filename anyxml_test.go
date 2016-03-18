package anyxml

import (
	"encoding/json"
	"fmt"
	"strings"
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
        "someotherkey": "someothervalue",
        "and": { "another": "k:v pair" },
        "list":[ "a string", 3.14159625, true ]

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
	if err != nil {
		t.Fatal(err)
	}
	x, err := Xml(i)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("[]->x:", string(x))

	a := []interface{}{"try", "this", 3.14159265, true}
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
	if err != nil {
		t.Fatal(err)
	}
	x, err := XmlIndent(i, "", "  ")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("[]->x:\n", string(x))

	a := []interface{}{"try", "this", 3.14159265, true}
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

type testEncodingExample struct {
	Key1 string
	Key2 string
	Key3 string
	Key4 string
}

func TestXMLEncoding(t *testing.T) {

	testEncodingSourceData := testEncodingExample{"I am fat & happy", "1 < 2", "2 > 1", "I'll have an apostrophe"}
	x, err := Xml(testEncodingSourceData)
	if err != nil {
		t.Fatal(err)
	}

	xmlString := string(x)
	fmt.Println("TestXMLEncoding XML out \n", xmlString)

	if strings.Contains(xmlString, "I am fat &amp; happy") == false {
		t.Fatal("Failed to encode ampersand\n")
	}

	if strings.Contains(xmlString, "1 &lt; 2") != true {
		t.Fatal("Failed to encode less than\n")
	}

	if strings.Contains(xmlString, "2 &gt; 1") != true {
		t.Fatal("Failed to encode greater than\n")
	}

	if strings.Contains(xmlString, "I&#39;ll have an apostrophe") != true {
		t.Fatal("Failed to encode apostrophe\n")
	}

}

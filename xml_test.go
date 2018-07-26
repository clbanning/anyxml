package anyxml

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestAnyXmlHeader2(t *testing.T) {
	fmt.Println("\n----------------  xml_test.go ...")
}

var anydata2 = []byte(`{ "element":[
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
]}`)

func TestXml2(t *testing.T) {
	var i interface{}
	err := json.Unmarshal(anydata2, &i)
	x, err := Xml(i)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("[]->x:", string(x))

}

func TestXmlIndent2(t *testing.T) {
	var i interface{}
	err := json.Unmarshal(anydata2, &i)
	x, err := XmlIndent(i, "", "  ")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("[]->x:\n", string(x))

}

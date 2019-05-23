
package anyxml

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestElemTagHeader(t *testing.T) {
	fmt.Println("\n----------------  elemtag_test.go ...")
}

func TestElemTag(t *testing.T) {
	MissingElementTag("myTag")
	defer MissingElementTag("")

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

	a := []interface{}{ "try", "this", 3.14159265, true }
	x, err = Xml(a)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("a->x:", string(x))
}

func TestElemTagIndent(t *testing.T) {
	MissingElementTag("myIndentTag")
	defer MissingElementTag("")

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

	a := []interface{}{ "try", "this", 3.14159265, true }
	x, err = XmlIndent(a, "", "  ")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("a->x:\n", string(x))
}

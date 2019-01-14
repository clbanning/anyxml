package anyxml

import (
	"fmt"
	"testing"
)

var s = `"'<>&`

func TestEscapeChars(t *testing.T) {
	fmt.Println("\n================== TestEscapeChars")

	ss := escapeChars(s)

	if ss != `&quot;&apos;&lt;&gt;&amp;` {
		t.Fatal(s, ":", ss)
	}

	fmt.Println(" s:", s)
	fmt.Println("ss:", ss)
}

func TestXMLEscapeChars(t *testing.T) {
	fmt.Println("================== TestXMLEscapeChars")

	XMLEscapeChars(true)
	defer XMLEscapeChars(false)

	m := map[string]interface{}{"mychars": s}

	x, err := XmlIndent(s, "", "  ")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("s:", string(x))

	x, err = XmlIndent(m, "", "  ")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("m:", string(x))
}

func TestXMLEscapeChars2(t *testing.T) {
	fmt.Println("================== TestXMLEscapeChars2")

	XMLEscapeChars(true)
	defer XMLEscapeChars(false)

	m := map[string]interface{}{"doc": map[string]interface{}{"simple": map[string]interface{}{"-attr1": "an attribute", "#text": `"'<>&`}}}
	fmt.Printf("m: %v\n", m)

	x, err := XmlIndent(m, "", "  ")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("x:", string(x))
}

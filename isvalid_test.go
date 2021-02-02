package anyxml

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestXmlCheckIsValid(t *testing.T) {
	fmt.Println("================== TestXmlCheckIsValid")

	data := []byte(`{"":"empty", "$invalid":"hex$", "entities":"<>&", "nil": null}`)
	m := make(map[string]interface{})
	err := json.Unmarshal(data, &m)
	if err != nil {
		t.Fatal("json.Unmarshal err;", err)
	}
	fmt.Printf("%v\n", m)

	XmlCheckIsValid()
	defer XmlCheckIsValid()
	if _, err := Xml(m); err == nil {
		t.Fatal("Xml err: nil")
	}

	if _, err := XmlIndent(m, "", "   "); err == nil {
		t.Fatal("XmlIndent err: nil")
	}

	ms := map[string]interface{}{
		"one":1,
		"not":"another",
	}
	fmt.Printf("%v\n", ms)
	if _, err = Xml(ms); err != nil {
		t.Fatal("Xml(ms) err:", err)
	}

	if _, err = XmlIndent(ms, "", "   "); err != nil {
		t.Fatal("XmlIndent(ms) err:", err)
	}
}

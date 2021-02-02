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
	if _, err := Xml(m); err == nil {
		t.Fatal("Xml err: nil")
	}

	if _, err := XmlIndent(m, "", "   "); err == nil {
		t.Fatal("XmlIndent err: nil")
	}
}

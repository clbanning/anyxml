// xml3_test.go - patch tests

package anyxml

import (
	"fmt"
	"testing"
)

func TestXml3(t *testing.T) {
	fmt.Println("\n------------ xml3_test.go")
}

// for: https://github.com/clbanning/mxj/pull/26
func TestOnlyAttributes(t *testing.T) {
	fmt.Println("========== TestOnlyAttributes")
	a := map[string]interface{}{ "-a":"try", "-b":"this", "-c":-3.14159265, "-d":true }
	x, err := XmlIndent(a, "", "  ")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("a->x:\n", string(x))
}


package main

import (
	"fmt"

	any "github.com/clbanning/anyxml"
)

func main() {
	data := map[interface{}]interface{}{
		"hello": "out there",
		1:       "number one",
		3.12:    "pi",
		"five":  5,
	}

	m, err := any.XmlIndent(data,"", "   ")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(m))
}

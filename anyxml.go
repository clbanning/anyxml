// anyxml - marshal an XML document from almost any variable
// 
// Wraps xml.Marshal with functionality in github.com/clbanning/mxj to create
// a more genericized XML marshaling capability. 
//
// See mxj package documentation for more information.  See anyxml_test.go for
// examples or just try Xml() or XmlIndent().
package anyxml

import (
	"encoding/xml"
	"reflect"
)

// Encode arbitrary value as XML.  Note: there are no guarantees.
func Xml(v interface{}, rootTag ...string) ([]byte, error) {
	if reflect.TypeOf(v).Kind() == reflect.Struct {
		return xml.Marshal(v)
	}

	var err error
	s := new(string)
	p := new(pretty)

	var rt string
	if len(rootTag) == 1 {
		rt = rootTag[0]
	} else {
		rt = DefaultRootTag
	}

	var ss string
	var b []byte
	switch v.(type) {
	case []interface{}:
		ss = "<" + rt + ">"
		for _, vv := range v.([]interface{}) {
			switch vv.(type) {
			case map[string]interface{}:
				m := vv.(map[string]interface{})
				if len(m) == 1 {
					for tag, val := range m {
						err = mapToXmlIndent(false, s, tag, val, p)
					}
				} else {
					err = mapToXmlIndent(false, s, rt, vv, p)
				}
			default:
				err = mapToXmlIndent(false, s, rt, vv, p)
			}
			if err != nil {
				break
			}
		}
		ss += *s + "</" + rt + ">"
		b = []byte(ss)
	case map[string]interface{}:
		b, err = anyxml(v.(map[string]interface{}),rootTag...)
	default:
		err = mapToXmlIndent(false, s, rt, v, p)
		b = []byte(*s)
	}

	return b, err
}


// Encode an arbitrary value as a pretty XML string. Note: there are no guarantees.
func XmlIndent(v interface{}, prefix, indent string, rootTag ...string) ([]byte, error) {
	if reflect.TypeOf(v).Kind() == reflect.Struct {
		return xml.MarshalIndent(v, prefix, indent)
	}

	var err error
	s := new(string)
	p := new(pretty)
	p.indent = indent
	p.padding = prefix

	var rt string
	if len(rootTag) == 1 {
		rt = rootTag[0]
	} else {
		rt = DefaultRootTag
	}

	var ss string
	var b []byte
	switch v.(type) {
	case []interface{}:
		ss = "<" + rt + ">\n"
		p.Indent()
		for _, vv := range v.([]interface{}) {
			switch vv.(type) {
			case map[string]interface{}:
				m := vv.(map[string]interface{})
				if len(m) == 1 {
					for tag, val := range m {
						err = mapToXmlIndent(true, s, tag, val, p)
					}
				} else {
					err = mapToXmlIndent(true, s, rt, vv, p)
				}
			default:
				err = mapToXmlIndent(true, s, rt, vv, p)
			}
			if err != nil {
				break
			}
		}
		ss += *s + "</" + rt + ">"
		b = []byte(ss)
	case map[string]interface{}:
		b, err = anyxmlIndent(v.(map[string]interface{}),prefix, indent, rootTag...)
	default:
		err = mapToXmlIndent(true, s, rt, v, p)
		b = []byte(*s)
	}

	return b, err
}

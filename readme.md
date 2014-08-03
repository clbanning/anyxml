
<h2>anyxml - create an XML document from almost any Go type</h2>
Marshal XML from map[string]interface{}, arrays, slices, and alpha/numeric values.  

This wraps encoding/xml with github.com/clbanning/mxj functionality.
See mxj package documentation for caveats, etc.

<h4>XML encoding conventions</h4>

   - 'nil' Map values, which may represent 'null' JSON values, are encoded as '\<tag/\>'.
      NOTE: the operation is not symmetric as '\<tag/\>' elements are decoded as 'tag:""' Map values,
            which, then, encode in JSON as '"tag":""' values..
   - in map[string]interface{} values keys that are prepended by a hyphen, '-', are assumed to be
     attributes.

<h4>Documentation</h4>

http://godoc.org/github.com/clbanning/anyxml

<h4>Example</h4>

Encode an arbitrary JSON object.<br>
<code>
	jasondata = []byte(`[<br>
		{ "somekey":"somevalue" },<br>
		"string",<br>
		3.14159265,<br>
		true<br>
	]`)<br>
	var i interface{}<br>
	err := json.Unmarshal(jsondaa, &i)<br>
	if err != nil {<br>
		// do something<br>
	}<br>
	x, err := anyxml.XmlIndent(i, "", "  ", "mydoc")<br>
	if err != nil {<br>
		// do something else<br>
	}<br>
	fmt.Println(string(x))<br>
// output:<br>
<mydoc><br>
	<somekey>somevalue</somekey><br>
	<element>string</element><br>
	<element>3.14159265</element><br>
	<element>true</true><br>
</mydoc><br>
</code>

See, also, xnyxml_test.go

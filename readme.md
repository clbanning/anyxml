
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


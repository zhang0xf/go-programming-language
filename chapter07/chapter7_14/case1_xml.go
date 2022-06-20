// Xmlselect prints the text of selected elements of an XML document.
package chapter7_14

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"strings"
)

// usage : 先编译chapter1_5.FetchUrl()
//         ./exercise http://www.w3.org/TR/2006/REC-xml11-20060816 > ./data/test.xml
//         再编译chapter7_14.Xml()
//         ./exercise < ./data/test.xml

// 四个主要的标记类型－StartElement，EndElement，CharData，和Comment－每一个都是encoding/xml包中的具体类型。
// 每一个对(*xml.Decoder).Token的调用都返回一个标记。

func Xml() {
	dec := xml.NewDecoder(os.Stdin)
	var stack []string // stack of element names
	for {
		tok, err := dec.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Fprintf(os.Stderr, "xmlselect: %v\n", err)
			os.Exit(1)
		}
		switch tok := tok.(type) {
		case xml.StartElement:
			stack = append(stack, tok.Name.Local) // push
		case xml.EndElement:
			stack = stack[:len(stack)-1] // pop
		case xml.CharData:
			if containsAll(stack, os.Args[1:]) {
				fmt.Printf("%s: %s\n", strings.Join(stack, " "), tok)
			}
		}
	}
}

// containsAll reports whether x contains the elements of y, in order.
func containsAll(x, y []string) bool {
	for len(y) <= len(x) {
		if len(y) == 0 {
			return true
		}
		if x[0] == y[0] {
			y = y[1:]
		}
		x = x[1:]
	}
	return false
}

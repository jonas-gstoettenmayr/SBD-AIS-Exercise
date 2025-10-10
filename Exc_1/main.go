package main

import (
	"fmt"
	"strings"
)

const TEXT = "Hello World!";

type Writer struct{
	text []string
}

// func (w *Writer) prnt(){ // can have lots of text so better not copy
// 	for _, v := range w.text{
// 		fmt.Println(v)
// 	}
// }

func (w Writer) String() string{
	return strings.Join(w.text, "\n")	
	
}

func main() {
	var content  = []string{TEXT, "Goodbye cruel World!"} // initialise array
	content = append(content, "THE END")
	writer := Writer{content}
	// (&writer).prnt() // with pointer take reference here
	fmt.Println(writer) // impliment String() to stringify and allow printing of types
}

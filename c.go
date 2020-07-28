package main

/*
	Go functions to be exported for use in C code. Note that if you use //export in a go file then you that preamble
	can only contain declarations and not definitions since the code generated from this file will be included in
	multiple files.
*/

import "C"
import "fmt"

//export GoOutput
func GoOutput(output *C.char) {
	fmt.Println(C.GoString(output))
}

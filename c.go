package main

import "C"
import "fmt"

//export GoOutput
func GoOutput(output *C.char) {
	fmt.Println(C.GoString(output))
}

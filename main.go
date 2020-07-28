package main

/*
Example of calling a C library from Go and vice-versa. The J interpreter connects to a front-end through user supplied
callbacks. This Go program calls the J shared library functions to initialize the interpreter, then starts a repl. The J
interpreter passes the output to the Go function `GoOutput`.
*/

// #cgo CFLAGS: -I./include
// #cgo LDFLAGS: -L. -Wl,-rpath=./ -lj -ldl
// #include "j.h"
// extern J JInit();
// extern int JDo(J, C*);
// extern void JSM(J, void *callbacks[]);
// extern void GoOutput(char *);
// void noop_callback(void) {
// }
// void output_callback(J jt, int result_code, char *output) {
// 	GoOutput(output);
// }
// void set_callbacks(J jt) {
//	void *callbacks[5] = {output_callback, noop_callback, noop_callback, noop_callback, (void *)3};
//	JSM(jt, callbacks);
// }
import "C"
import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	jt := C.JInit()
	C.set_callbacks(jt)
	fmt.Print("> ")
	for scanner.Scan() {
		line := scanner.Text()
		C.JDo(jt, (*C.char)(C.CString(line)))
		fmt.Print("> ")
	}
}

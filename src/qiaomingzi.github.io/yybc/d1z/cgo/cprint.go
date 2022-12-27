package main

import "C"
import "unsafe"

func main() {
	cstr := C.String("hello")
	C.puts(cstr)
	C.free(unsafe.Pointer(cstr))
}

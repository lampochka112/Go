package main

// #cgo LDFLAGS: -L. -lcpplib
// #include "lib.h"
import "C"

func main() {
    C.my_cpp_function(C.CString("Gopher"))
}
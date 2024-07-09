package main

import "C"

func main() {
	// Nothing here
}

//export Foo
func Foo(val C.ulong) C.ulong {
	return val + 1
}

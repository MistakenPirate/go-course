package main

import "fmt"

func main() {
	var ptr *int
	var a int = 10
	ptr = &a

	fmt.Println(a)
	fmt.Println(&a)

	fmt.Println(ptr)
	fmt.Println(*ptr) // dereferencing a pointer

	// zero value of a pointer is nil
	// if ptr == nil {
	// 	fmt.Println("Pointer in nil")
	// }

	modifyValue(ptr)
	fmt.Println(a)
}

func modifyValue(ptr *int) {
	*ptr++
}

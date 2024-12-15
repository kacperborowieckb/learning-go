package main

import "fmt"

func main() {
	// pointer is a special type
	// they hold a memory address of a variable
	// var p = 0x0002...

	// pointer to int32, for now nil because it doesn't point to anything
	// var p *int32

	// to make it hold some memory address we can use new()
	var p *int32 = new(int32)
	fmt.Println("P address:", p)
	fmt.Println("P value:", *p)
	// to change a value of a pointer value
	*p = 15
	fmt.Println("P value:", *p)

	// normal int32
	var i int32
	fmt.Println("Default i:", i)

	p = &i

	// slices by copy to a function or overall variables
	// in normal flow it will copy the address so it will modify both values
	// we can pass pointer so some e.g slice to a function to safe some space
	// and modify the slice in the main/parent(?) function

}

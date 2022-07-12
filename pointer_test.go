package gu

import (
	"fmt"
	"testing"
)

func TestPtr(t *testing.T) {
	sp := Ptr("Hello world!")
	if sp == nil {
		t.Fatal("Ptr returned nil")
	}
}

func ExamplePtr() {
	// Pointer of a string
	// stringPointer := &"Hello world!": invalid operation: cannot take address of "Hello world!" (untyped string constant)
	stringPointer := Ptr("Hello world!")
	fmt.Printf("stringPointer is of type %T and points to value %v\n", stringPointer, *stringPointer)

	// Constant
	const i int64 = 22

	// int64Pointer := &i: invalid operation: cannot take address of i (constant 22 of type int64)
	int64Pointer := Ptr(i)
	fmt.Printf("int64Pointer is of type %T and points to value %v\n", int64Pointer, *int64Pointer)

	// Function return
	// funcReturn := &fmt.Sprint(99): invalid operation: cannot take address of fmt.Sprint(99) (value of type string)
	funcReturn := Ptr(fmt.Sprint(99))
	fmt.Printf("funcReturn is of type %T and points to value %v\n", funcReturn, *funcReturn)

	// Output: stringPointer is of type *string and points to value Hello world!
	// int64Pointer is of type *int64 and points to value 22
	// funcReturn is of type *string and points to value 99
}

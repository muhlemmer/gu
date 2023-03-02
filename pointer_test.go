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

func TestValue(t *testing.T) {
	tests := []struct {
		pointer   *string
		wantValue string
	}{
		{
			nil,
			"",
		},
		{
			Ptr("foo"),
			"foo",
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.pointer), func(t *testing.T) {
			if gotValue := Value(tt.pointer); gotValue != tt.wantValue {
				t.Errorf("Value() = %v, want %v", gotValue, tt.wantValue)
			}
		})
	}
}

func ExampleValue() {
	type document struct {
		ID          *int    `json:"id,omitempty"`
		Description *string `json:"description,omitempty"`
	}

	d := document{
		Description: Ptr("foobar"),
	}

	// this would panic, d.ID is nil
	// fmt.Println(*d.ID, *d.Description)

	// d.ID is nil, so a 0 is printed
	fmt.Println(Value(d.ID), Value(d.Description))

	// Output: 0 foobar
}

func TestPtrCopy(t *testing.T) {
	type args struct {
		pointer *int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "nil",
			args: args{},
		},
		{
			name: "value",
			args: args{
				Ptr(1),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PtrCopy(tt.args.pointer); got == tt.args.pointer {
				t.Errorf("PtrCopy(): %v == %v", got, tt.args.pointer)
			}
		})
	}
}

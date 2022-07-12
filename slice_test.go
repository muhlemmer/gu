package gu

import (
	"fmt"
	"log"
	"reflect"
	"strings"
	"testing"
)

func TestInterfaceSlice(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5}
	want := []interface{}{1, 2, 3, 4, 5}

	got := InterfaceSlice(slice)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("InterfaceSlice =\n%v\nwant\n%v", got, want)
	}
}

func ExampleInterfaceSlice() {
	stringSlice := []string{"Hello", ", ", "World", "!"}
	// fmt.Print(stringSlice...): cannot use stringSlice (variable of type []string) as type []any in argument to fmt.Print
	fmt.Print(InterfaceSlice(stringSlice)...)

	// Output: Hello, World!
}

func TestAssertInterfaces(t *testing.T) {
	tests := []struct {
		name    string
		is      []interface{}
		want    []int
		wantErr bool
	}{
		{
			"success",
			[]interface{}{1, 2, 3, 4, 5},
			[]int{1, 2, 3, 4, 5},
			false,
		},
		{
			"error",
			[]interface{}{1, 2, 3, 4, "foo"},
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := AssertInterfaces[int](tt.is)
			if (err != nil) != tt.wantErr {
				t.Errorf("AssertInterfaces() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AssertInterfaces() = %v, want %v", got, tt.want)
			}
		})
	}
}

func ExampleAssertInterfaces() {
	interfaceSlice := []interface{}{"Hello", "World!"}
	stringSlice, err := AssertInterfaces[string](interfaceSlice)
	if err != nil {
		log.Fatal(err)
	}

	s := strings.Join(stringSlice, ", ")
	fmt.Println(s)

	interfaceSlice = []interface{}{1, 1.1, "foobar"}
	intSlice, err := AssertInterfaces[int](interfaceSlice)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(intSlice)

	// Output: Hello, World!
	// cannot assert float64 of value 1.1 to int at index 1
	// []
}

func TestAssertInterfacesP(t *testing.T) {
	tests := []struct {
		name    string
		is      []interface{}
		want    []int
		wantErr bool
	}{
		{
			"success",
			[]interface{}{1, 2, 3, 4, 5},
			[]int{1, 2, 3, 4, 5},
			false,
		},
		{
			"error",
			[]interface{}{1, 2, 3, 4, "foo"},
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				err, _ := recover().(error)
				if (err != nil) != tt.wantErr {
					t.Errorf("AssertInterfaces() error = %v, wantErr %v", err, tt.wantErr)
				}
			}()

			got := AssertInterfacesP[int](tt.is)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AssertInterfaces() = %v, want %v", got, tt.want)
			}
		})
	}
}

func ExampleAssertInterfacesP() {
	interfaceSlice := []interface{}{"Hello", "World!"}

	s := strings.Join(AssertInterfacesP[string](interfaceSlice), ", ")
	fmt.Println(s)

	// Output: Hello, World!
}

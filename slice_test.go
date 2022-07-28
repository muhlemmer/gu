package gu

import (
	"fmt"
	"log"
	"reflect"
	"strconv"
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

func TestTransform(t *testing.T) {
	type testA struct {
		ID int32
		S  []string
	}

	type testB struct {
		ID int64
		S  string
	}

	transFunc := func(a testA) testB {
		return testB{
			ID: int64(a.ID),
			S:  strings.Join(a.S, ", "),
		}
	}

	tests := []struct {
		name string
		as   []testA
		want []testB
	}{
		{
			"nil",
			nil,
			nil,
		},
		{
			"entries",
			[]testA{
				{
					ID: 1,
					S:  []string{"Hello", "World!"},
				},
				{
					ID: 2,
					S:  []string{"foo", "bar"},
				},
				{
					ID: 3,
					S:  []string{"spanac"},
				},
			},
			[]testB{
				{
					ID: 1,
					S:  "Hello, World!",
				},
				{
					ID: 2,
					S:  "foo, bar",
				},
				{
					ID: 3,
					S:  "spanac",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Transform(tt.as, transFunc); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Transform() = %v, want %v", got, tt.want)
			}
		})
	}
}

func ExampleTransform_itoa() {
	in := []int{1, 2, 3, 4, 5}

	out := Transform(in, strconv.Itoa)
	fmt.Printf("out is of type %T and contains %v", out, out)

	// Output: out is of type []string and contains [1 2 3 4 5]
}

func ExampleTransform_struct() {
	type A struct {
		ID int32
		S  []string
	}

	type B struct {
		ID int64
		S  string
	}

	// define a tranformer function
	transFunc := func(a A) B {
		return B{
			ID: int64(a.ID),
			S:  strings.Join(a.S, ", "),
		}
	}

	in := []A{
		{
			ID: 1,
			S:  []string{"Hello", "World!"},
		},
		{
			ID: 2,
			S:  []string{"foo", "bar"},
		},
		{
			ID: 3,
			S:  []string{"spanac"},
		},
	}

	// create the transformed slice
	out := Transform(in, transFunc)

	fmt.Printf("out is of type %T and contains %v", out, out)
	// Output: out is of type []gu.B and contains [{1 Hello, World!} {2 foo, bar} {3 spanac}]
}

func TestTransformErr(t *testing.T) {
	tests := []struct {
		name    string
		as      []string
		want    []int
		wantErr bool
	}{
		{
			"nil",
			nil,
			nil,
			false,
		},
		{
			"succes",
			[]string{"1", "2", "3", "4", "5"},
			[]int{1, 2, 3, 4, 5},
			false,
		},
		{
			"error",
			[]string{"1", "2", "foo", "4", "5"},
			[]int{1, 2},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := TransformErr(tt.as, strconv.Atoi)
			if (err != nil) != tt.wantErr {
				t.Errorf("TransformErr() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TransformErr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func ExampleTransformErr_atoi() {
	in := []string{"1", "2", "3", "4", "5"}

	out, err := TransformErr(in, strconv.Atoi)
	if err != nil {
		panic(err)
	}

	fmt.Printf("out is of type %T and contains %v\n", out, out)

	// this will cause an error
	in = []string{"1", "2", "foo", "4", "5"}

	out, err = TransformErr(in, strconv.Atoi)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("out is of type %T and contains %v\n", out, out)

	// Output: out is of type []int and contains [1 2 3 4 5]
	// transform index 2: strconv.Atoi: parsing "foo": invalid syntax
	// out is of type []int and contains [1 2]
}

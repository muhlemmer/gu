package gu

import "fmt"

// InterfaceSlice transforms a slice of any type to a slice of interface{}.
// This can be usefull when you have a slice of concrete types that has to be passed
// to a function that takes a (variadic) slice of interface{}.
func InterfaceSlice[T any](slice []T) []interface{} {
	out := make([]interface{}, len(slice))

	for i := 0; i < len(slice); i++ {
		out[i] = slice[i]
	}

	return out
}

// AssertInterfaces asserts all members of the passed slice of interfaces
// to the requested type and returs a slice of that type.
// A nil slice allong with an error is returned
// if one of the entries cannot be asserted to the destination type.
func AssertInterfaces[T any](is []interface{}) ([]T, error) {
	out := make([]T, len(is))

	for i := 0; i < len(is); i++ {
		var ok bool

		if out[i], ok = is[i].(T); !ok {
			return nil, fmt.Errorf("cannot assert %T of value %v to %T at index %d", is[i], is[i], out[i], i)
		}
	}

	return out, nil
}

// AssertInterfacesP is like AssertInterfaces, only that it does not
// check for succesfull assertion and lets the runtime panic if assertion fails.
// Usefull for inlining in cases where you are 100% sure of the concrete type of the passed slice.
func AssertInterfacesP[T any](is []interface{}) []T {
	out := make([]T, len(is))

	for i := 0; i < len(is); i++ {
		out[i] = is[i].(T)
	}

	return out
}

// A transformer copies values into a new instance of 'B'.
// How values are copied is up to the implementation.
type Transformer[B any] interface {
	Transform() B
}

// Transform a slice of type 'A' to a slice of type 'B'.
// 'A' must implement the Tranformer interface,
// which returns a single instance of 'B'.
//
// Usefull when working with slices of different, but similar,
// struct types and you don't want to write the 'for' loops
// over and over again.
func Transform[B any, A Transformer[B]](as []A) []B {
	if as == nil {
		return nil
	}

	out := make([]B, len(as))

	for i, a := range as {
		out[i] = a.Transform()
	}

	return out
}

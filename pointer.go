package gu

// Ptr returns a pointer to the passed value.
// This helps where the ampersand can't be used directly,
// such as on constants or function returns.
func Ptr[T any](value T) (pointer *T) {
	return &value
}

// Value return the value behind pointer.
// If the pointer is nil, the zero / empty value of T is returned.
// This helps to safely access variables where it does not matter of the program
// if they where nil or not, but you want to prevent a panic.
//
// Common use case is fields in gerated structs from frameworks such as
// protobuf 2 or openapi 3.
func Value[T any](pointer *T) (value T) {
	if pointer != nil {
		value = *pointer
	}

	return value
}

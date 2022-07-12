package gu

// Ptr returns a pointer to the passed value.
// This helps where the ampersand can't be used directly,
// such as on constants or function returns.
func Ptr[T any](v T) *T {
	return &v
}

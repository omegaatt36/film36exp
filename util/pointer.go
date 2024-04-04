package util

// Pointer returns a pointer to a value
func Pointer[T any](v T) *T {
	return &v
}

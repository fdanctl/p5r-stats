package utils

// FallbackToB returns the second argument if the first is nil, otherwise it returns the first argument.
func FallbackToB[T any](a, b *T) T {
	if a == nil {
		return *b
	}
	return *a
}

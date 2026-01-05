package utils

// if a is nil fallback to b
func FallbackToB[T any](a, b *T) T {
	if a == nil {
		return *b
	}
	return *a
}

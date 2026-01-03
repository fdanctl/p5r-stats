package utils

// if a is nil fallback to b
func FallbackToB[T comparable](a, b *T) T {
	if a == nil {
		return *b
	}
	return *a
}

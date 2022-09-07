package gu

// MapEqual check if two maps have exactly the same content.
// If both maps are nil, they are considered equal.
// When a nil map is compared to an empty map,
// they are not considered equal.
func MapEqual[K, V comparable](a, b map[K]V) bool {
	if a == nil && b == nil {
		return true
	}

	if (a == nil && b != nil) || (b == nil && a != nil) {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for k, av := range a {
		if bv, ok := b[k]; !ok || av != bv {
			return false
		}
	}

	return true
}


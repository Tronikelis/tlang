package slice

func SafeIndex[G any, T []G](s T, index int) *G {
	if index < 0 {
		return nil
	}

	if len(s)-1 < index {
		return nil
	}

	return &s[index]
}

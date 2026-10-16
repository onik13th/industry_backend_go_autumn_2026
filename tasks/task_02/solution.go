package main

func rotateRunes(s string, shift int) string {
	runes := []rune(s)
	if len(runes) == 0 {
		return ""
	}

	shift %= len(runes)
	if shift < 0 {
		shift += len(runes)
	}

	result := make([]rune, len(runes))
	copied := copy(result, runes[shift:])
	copy(result[copied:], runes[:shift])

	return string(result)
}

package palindrom

// Is checks that the texts first half mirrors the rest.
// A real palindrom needs to be a real word, this is not checked.
func Is(text string) bool {
	runes := []rune(text)
	length := len(runes)

	for i := 0; i < length/2; i++ {
		if runes[i] != runes[length-1-i] {
			return false
		}
	}

	return true
}

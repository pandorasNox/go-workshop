package pangram

import (
	"unicode"
)

// Is checks that every alphabetical charactar is contained in the text.
func Is(text string) bool {
	alphabet := make(map[rune]bool)
	for i := 'a'; i <= 'z'; i++ {
		alphabet[i] = true
	}

	for _, c := range text {
		delete(alphabet, unicode.ToLower(c))
	}

	return len(alphabet) == 0
}

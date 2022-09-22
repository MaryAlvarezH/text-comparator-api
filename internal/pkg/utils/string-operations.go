package utils

import (
	"unicode"

	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

// GetVowelsNumber returns the number of vowels in a given string
func GetVowelsNumber(str string) int {
	s := NormalizeText(str)
	c := 0
	for _, value := range s {
		switch value {
		case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
			c++
		}
	}

	return c
}

// NormalizeText applies text normalization to remove special signs in strings
// NFD (descomposición canónica de formato de normalización) - Los caracteres se descomponen según su equivalencia canónica.
// NFD (composición canónica de formato de normalización) - Los caracteres se descomponen y después se recomponen según su equivalencia canónica..
func NormalizeText(str string) string {
	t := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
	s, _, _ := transform.String(t, str)
	return s
}

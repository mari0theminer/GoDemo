package string

import "strings"

func makeCamelCase(s string) string {
	var newString string

	words := strings.Fields(s)
	for _, word := range words {
		newString = newString + makeFirstLetterBig(word)
	}
	return newString
}
func makeKebabCase(s string) string {
	stringWithMinus := strings.ReplaceAll(s, " ", "-")
	return strings.ToLower(stringWithMinus)

}
func makeFirstLetterBig(s string) string {
	var first bool = true
	var newString string

	for _, char := range s {
		if first {
			first = false
			newString = newString + strings.ToUpper(string(char))
		} else {
			newString = newString + strings.ToLower(string(char))

		}
	}
	return newString
}

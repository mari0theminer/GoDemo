package string

import "strings"

func MakeCamelCase(s string) string {
	var newString string

	words := strings.Fields(s)
	for _, word := range words {
		newString = newString + MakeFirstLetterBig(word)
	}
	return newString
}
func MakeKebabCase(s string) string {
	stringWithMinus := strings.ReplaceAll(s, " ", "-")
	return strings.ToLower(stringWithMinus)

}
func MakeFirstLetterBig(s string) string {
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

func ReplaceNyan(s string) string {
	return "Nyan cat"
}

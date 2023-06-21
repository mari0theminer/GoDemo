package sllicy

func Push(s *[]string, new string) {
	*s = append(*s, new)
}

func Pop(s *[]string) string {
	defer func() { *s = (*s)[:len(*s)-1] }()
	return (*s)[len(*s)-1]
}

func Peek(s []string) string {
	return s[len(s)-1]

}

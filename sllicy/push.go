package sllicy

func push(s *[]string, new string) {
	*s = append(*s, new)
}

func pop(s *[]string) string {
	defer func() { *s = (*s)[:len(*s)-1] }()
	return (*s)[len(*s)-1]
}

func peek(s []string) string {
	return s[len(s)-1]

}

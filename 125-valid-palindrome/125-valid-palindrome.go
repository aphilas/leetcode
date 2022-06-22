func isPalindrome(s string) bool {
    re := regexp.MustCompile(`[^a-zA-Z0-9]`)
	s = strings.ToLower(re.ReplaceAllLiteralString(s, ""))

	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}

	return true
}
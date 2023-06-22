package asciiart

func Index(s string, t string) int {
	len1, len2 := len(s), len(t)
	for i := 0; i <= len1-len2; i++ {
		if s[i:i+len2] == t {
			return i
		}
	}
	return -1
}

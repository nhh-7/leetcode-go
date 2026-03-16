package times1

func longestPalindrome(s string) string {
	start := 0
	maxLen := 0
	n := len(s)
	for i := range s {
		l, r := i, i
		for l >= 0 && r < n && s[l] == s[r] {
			l--
			r++
		}
		if r-l-1 > maxLen {
			start = l + 1
			maxLen = r - l - 1
		}
	}
	for i := range s {
		l, r := i, i+1
		for l >= 0 && r < n && s[l] == s[r] {
			l--
			r++
		}
		if r-l-1 > maxLen {
			start = l + 1
			maxLen = r - l - 1
		}
	}
	return s[start : start+maxLen]
}

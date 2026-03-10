package internal

func longestPalindrome(s string) string {
	n := len(s)
	maxLen, start := 0, 0
	for i := range n {
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
	for i := range n - 1 {
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

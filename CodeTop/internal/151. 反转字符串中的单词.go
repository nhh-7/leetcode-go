package internal

import "slices"

func reverseWords(s string) string {
	b := []byte(s)
	n := len(s)

	slow, fast := 0, 0
	for fast < n {
		for fast < n && b[fast] == ' ' {
			fast++
		}
		if fast < n {
			if slow != 0 {
				b[slow] = ' '
				slow++
			}
			for fast < n && b[fast] != ' ' {
				b[slow] = b[fast]
				fast++
				slow++
			}
		}
	}
	b = b[:slow]

	slices.Reverse(b)

	start := 0
	for i := 0; i <= len(b); i++ {
		if i == len(b) || b[i] == ' ' {
			slices.Reverse(b[start:i])
			start = i + 1
		}
	}
	return string(b)
}

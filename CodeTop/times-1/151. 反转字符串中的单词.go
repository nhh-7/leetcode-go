package times1

import "slices"

func reverseWords(s string) string {
	slow, fast := 0, 0
	n := len(s)

	str := []byte(s)
	for fast < n {
		for fast < n && str[fast] == ' ' {
			fast++
		}

		if fast < n {
			if slow > 0 {
				str[slow] = ' '
				slow++
			}

			for fast < n && str[fast] != ' ' {
				str[slow] = str[fast]
				fast++
				slow++
			}
		}
	}
	str = str[:slow]
	slices.Reverse(str)

	start := 0
	for i := 0; i <= len(str); i++ {
		if i == len(str) || str[i] == ' ' {
			slices.Reverse(str[start:i])
			start = i + 1
		}
	}
	return string(str)
}

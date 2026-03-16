package internal

import (
	"strings"
	"unicode"
)

func decodeString(s string) string {
	type pair struct {
		k   int
		str string
	}
	stack := []pair{}

	k := 0
	res := ""
	for _, ch := range s {
		if unicode.IsDigit(ch) {
			k = k*10 + int(ch-'0')
		} else if unicode.IsLetter(ch) {
			res += string(ch)
		} else if ch == '[' {
			stack = append(stack, pair{k, res})
			k = 0
			res = ""
		} else if ch == ']' {
			pre := stack[len(stack)-1].str
			times := stack[len(stack)-1].k
			stack = stack[:len(stack)-1]
			res = pre + strings.Repeat(res, times)
		}
	}
	return res
}

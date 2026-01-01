package stack

import (
	"strings"
	"unicode"
)

/*
遇到[就压入栈，遇到]读取栈顶进行处理
*/
func decodeString(s string) string {
	type pair struct {
		k int
		s string
	}

	stack := []pair{}
	res := ""
	k := 0
	for _, c := range s {
		if unicode.IsLetter(c) {
			res += string(c)
		} else if unicode.IsDigit(c) {
			k = k*10 + int(c-'0')
		} else if c == '[' {
			stack = append(stack, pair{k, res})
			// res为将拼接到到这次重复字符串的前缀字符串
			res = ""
			k = 0
		} else {
			p := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			pre := p.s
			// 重复k次拼接前缀，获得下一次碰到[的前缀
			// res记录了在[]中的需要重复的字符串
			res = pre + strings.Repeat(res, p.k)
		}
	}
	return res
}

package main

import (
	"strconv"
	"strings"
)

func compressString(S string) string {
	if len(S) == 0 {
		return ""
	}
	n := len(S)
	count := 1
	var sb strings.Builder

	for i := 1; i < n; i++ {
		if S[i] == S[i-1] {
			count++
		} else {
			sb.WriteByte(S[i-1])
			sb.WriteString(strconv.Itoa(count))
			count = 1
		}
	}

	sb.WriteByte(S[n-1])
	sb.WriteString(strconv.Itoa(count))

	res := sb.String()

	if len(res) >= len(S) {
		return S
	}
	return res
}

package times2

func myAtoi(s string) int {
	n := len(s)
	idx := 0
	for idx < n && s[idx] == ' ' {
		idx++
	}

	sign := 1
	if idx < n && s[idx] == '-' {
		sign = -1
		idx++
	} else if idx < n && s[idx] == '+' {
		idx++
	}

	maxInt := 1<<31 - 1
	minInt := -1 << 31

	ans := 0
	for ; idx < n && s[idx] >= '0' && s[idx] <= '9'; idx++ {
		digit := int(s[idx] - '0')

		if ans > maxInt/10 || (ans == maxInt/10 && digit > 7) {
			if sign == 1 {
				return maxInt
			} else {
				return minInt
			}
		}
		ans = ans*10 + digit
	}
	return sign * ans
}

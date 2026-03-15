package internal

func myAtoi(s string) int {
	n := len(s)
	idx := 0

	for idx < n && s[idx] == ' ' {
		idx++
	}

	if idx == n {
		return 0
	}

	sign := 1
	if s[idx] == '+' {
		idx++
	} else if s[idx] == '-' {
		idx++
		sign = -1
	}

	max_int := 1<<31 - 1
	min_int := -1 << 31
	res := 0
	for idx < n && s[idx] >= '0' && s[idx] <= '9' {
		digit := int(s[idx] - '0')

		if res > max_int/10 || (res == max_int/10 && digit > 7) {
			if sign == 1 {
				return max_int
			} else {
				return min_int
			}
		}

		res = res*10 + digit
		idx++
	}
	return res * sign
}

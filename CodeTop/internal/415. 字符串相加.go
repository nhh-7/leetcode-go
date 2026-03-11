package internal

import "strconv"

func addStrings(num1 string, num2 string) string {
	i, j := len(num1)-1, len(num2)-1
	carry := 0

	res := ""
	for i >= 0 || j >= 0 || carry > 0 {
		x, y := 0, 0
		if i >= 0 {
			x = int(num1[i] - '0')
			i--
		}
		if j >= 0 {
			y = int(num2[j] - '0')
			j--
		}

		sum := x + y + carry
		res = strconv.Itoa(sum%10) + res
		carry = sum / 10
	}
	return res
}

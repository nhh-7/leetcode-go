package times2

import "strconv"

func addStrings(num1 string, num2 string) string {
	m, n := len(num1), len(num2)
	i, j, carry := m-1, n-1, 0
	res := ""
	for i >= 0 || j >= 0 || carry > 0 {
		val1, val2 := 0, 0
		if i >= 0 {
			val1 = int(num1[i] - '0')
			i--
		}
		if j >= 0 {
			val2 = int(num2[j] - '0')
			j--
		}

		sum := val1 + val2 + carry
		carry = sum / 10
		res = strconv.Itoa(sum%10) + res
	}
	return res
}

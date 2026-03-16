package times1

import "strconv"

func addStrings(num1 string, num2 string) string {

	p1, p2 := len(num1)-1, len(num2)-1
	carry := 0
	res := ""
	for p1 >= 0 || p2 >= 0 || carry > 0 {
		v1, v2 := 0, 0
		if p1 >= 0 {
			v1 = int(num1[p1] - '0')
			p1--
		}
		if p2 >= 0 {
			v2 = int(num2[p2] - '0')
			p2--
		}
		sum := v1 + v2 + carry
		carry = sum / 10
		res = strconv.Itoa(sum%10) + res
	}
	return res
}

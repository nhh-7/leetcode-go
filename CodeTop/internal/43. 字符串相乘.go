package internal

import "strconv"

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}

	m, n := len(num1), len(num2)

	pos := make([]int, m+n)
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			mul := int(num1[i]-'0') * int(num2[j]-'0')

			p1, p2 := i+j, i+j+1 // 分别是进位位 和 当前位

			sum := mul + pos[p2]

			pos[p2] = sum % 10
			pos[p1] += sum / 10
		}
	}

	res := ""
	for _, v := range pos {
		// 如果还没遇到非零数字，且当前是 0，就跳过
		if len(res) == 0 && v == 0 {
			continue
		}
		res += strconv.Itoa(v)
	}
	return res
}

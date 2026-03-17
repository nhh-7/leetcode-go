package times1

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

			p1, p2 := i+j, i+j+1

			sum := mul + pos[p2]
			pos[p1] += sum / 10
			pos[p2] = sum % 10
		}
	}

	res := ""
	for _, v := range pos {
		if v == 0 && len(res) == 0 {
			continue
		}
		res = res + strconv.Itoa(v)
	}
	return res
}

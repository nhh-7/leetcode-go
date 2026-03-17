package times1

func compareVersion(version1 string, version2 string) int {
	n1, n2 := len(version1), len(version2)

	i, j := 0, 0
	for i < n1 || j < n2 {
		x := 0
		for ; i < n1 && version1[i] != '.'; i++ {
			x = x*10 + int(version1[i]-'0')
		}
		i++

		y := 0
		for ; j < n2 && version2[j] != '.'; j++ {
			y = y*10 + int(version2[j]-'0')
		}
		j++
		if x > y {
			return 1
		} else if x < y {
			return -1
		}
	}
	return 0
}

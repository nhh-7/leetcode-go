package times2

func compareVersion(version1 string, version2 string) int {
	m, n := len(version1), len(version2)

	i, j := 0, 0
	for i < m || j < n {
		x, y := 0, 0
		for ; i < m && version1[i] != '.'; i++ {
			x = x*10 + int(version1[i]-'0')
		}
		i++

		for ; j < n && version2[j] != '.'; j++ {
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

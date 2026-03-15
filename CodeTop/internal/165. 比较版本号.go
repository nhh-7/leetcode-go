package internal

import (
	"strconv"
	"strings"
)

func compareVersion(version1 string, version2 string) int {
	v1, v2 := strings.Split(version1, "."), strings.Split(version2, ".")

	n1, n2 := len(v1), len(v2)
	maxLen := max(n1, n2)

	for i := range maxLen {
		nums1, nums2 := 0, 0
		if i < n1 {
			nums1, _ = strconv.Atoi(v1[i])
		}
		if i < n2 {
			nums2, _ = strconv.Atoi(v2[i])
		}

		if nums1 < nums2 {
			return -1
		} else if nums1 > nums2 {
			return 1
		}
	}
	return 0
}

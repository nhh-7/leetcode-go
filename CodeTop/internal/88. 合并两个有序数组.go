package internal

func merge1(nums1 []int, m int, nums2 []int, n int) {
	p1, p2 := m-1, n-1
	pTail := m + n - 1

	for p2 >= 0 {
		if p1 >= 0 && nums1[p1] > nums2[p2] {
			nums1[pTail] = nums1[p1]
			p1--
		} else {
			nums1[pTail] = nums2[p2]
			p2--
		}
		pTail--
	}
}

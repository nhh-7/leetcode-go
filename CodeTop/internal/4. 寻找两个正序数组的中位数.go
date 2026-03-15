package internal

/*
在两个数组中找到一个“分割线”，使得左半部分的元素个数等于右半部分（或多一个），且左侧所有元素都小于右侧。

假设我们将两个数组 A 和 B 分别切一刀：
		A 在 i 处切开，左侧有 i 个元素。
		B 在 j 处切开，左侧有 j 个元素。
我们需要满足两个条件：
		个数平衡：i + j = (m + n + 1) / 2左右两边元素个数基本相等）。
		数值有序：A[i-1] <= B[j]（A 的左边最大值 < B 的右边最小值）
				B[j-1] <= A[i]（B 的左边最大值 < A 的右边最小值）

选择短数组进行二分：假设 nums1 比 nums2 短。我们只在 nums1 的范围 [0, m] 里进行二分查找，(i可以处于所有元素之后)，确定位置 i。
计算 j：根据个数平衡公式，直接算出 j = ((m+n+1）/ 2) - i。
判断与移动：如果 A[i-1] > B[j]，说明 i 太大了，需要左移（right = i - 1）。
		  如果 B[j-1] > A[i]，说明 i 太小了，需要右移（left = i + 1）。
边界处理：如果 i 或 j 到了数组尽头（0 或 m/n），我们将左侧最大值视为负无穷 ，右侧最小值视为正无穷。


*/

func findMedianSortedArrays1(nums1 []int, nums2 []int) float64 {
	// 确保 nums1 是短的，这样二分更快，且 j 不会是负数
	if len(nums1) > len(nums2) {
		return findMedianSortedArrays1(nums2, nums1)
	}

	m, n := len(nums1), len(nums2)
	left, right := 0, m
	halfLen := (m + n + 1) / 2

	for left <= right {
		i := (left + right) / 2
		j := halfLen - i

		if i < m && nums2[j-1] > nums1[i] {
			left = i + 1 // i 太小了
		} else if i > 0 && nums1[i-1] > nums2[j] {
			right = i - 1 // i 太大了
		} else {
			// 找到了完美的 i
			var maxLeft int
			if i == 0 {
				maxLeft = nums2[j-1]
			} else if j == 0 {
				maxLeft = nums1[i-1]
			} else {
				maxLeft = max(nums1[i-1], nums2[j-1])
			}

			if (m+n)%2 == 1 {
				return float64(maxLeft) // 奇数长度
			}

			var minRight int
			if i == m {
				minRight = nums2[j]
			} else if j == n {
				minRight = nums1[i]
			} else {
				minRight = min(nums1[i], nums2[j])
			}

			return float64(maxLeft+minRight) / 2.0 // 偶数长度
		}
	}
	return 0.0
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) > len(nums2) {
		return findMedianSortedArrays(nums2, nums1)
	}

	m, n := len(nums1), len(nums2)
	halfLen := (m + n + 1) / 2
	left, right := 0, m

	for left <= right {
		i := (left + right) / 2
		j := halfLen - i

		if i < m && nums2[j-1] > nums1[i] { // nums1的右侧最小值比nums2的左侧最大值小， i太小了
			left = i + 1
		} else if i > 0 && nums1[i-1] > nums2[j] { // nums1的左侧最大值比nums2的右侧最小值，i太大了
			right = i - 1
		} else {
			maxLeft := 0
			if i == 0 {
				maxLeft = nums2[j-1]
			} else if j == 0 {
				maxLeft = nums1[i-1]
			} else {
				maxLeft = max(nums1[i-1], nums2[j-1])
			}

			if (m+n)%2 == 1 {
				return float64(maxLeft)
			}

			minRight := 0
			if i == m {
				minRight = nums2[j]
			} else if j == n {
				minRight = nums1[i]
			} else {
				minRight = min(nums1[i], nums2[j])
			}
			return float64(maxLeft+minRight) / 2.0
		}
	}
	return 0.0
}

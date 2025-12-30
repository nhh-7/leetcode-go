package binarysearch

func searchMatrix1(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])

	// 先搜索第一列判断目标存在哪一行
	left, right := 0, m-1
	for left <= right {
		mid := (left + right) / 2
		if matrix[mid][0] > target {
			right = mid - 1
		} else if matrix[mid][0] < target {
			left = mid + 1
		} else {
			return true
		}
	}

	if right < 0 {
		return false
	}

	// 现在right为目标存在的那一行
	row := right

	// 搜寻目标行
	left, right = 0, n-1
	for left <= right {
		mid := (left + right) / 2
		if matrix[row][mid] > target {
			right = mid - 1
		} else if matrix[row][mid] < target {
			left = mid + 1
		} else {
			return true
		}
	}
	return false
}

// 直接将整个二维数组当作一个连续的一维升序数组
func searchMatrix(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])

	left, right := 0, m*n-1
	for left <= right {
		mid := left + (right-left)/2
		row, col := mid/n, mid%n
		if matrix[row][col] > target {
			right = mid - 1
		} else if matrix[row][col] < target {
			left = mid + 1
		} else {
			return true
		}
	}
	return false
}

package times2

func mySqrt(x int) int {
	left, right := 0, x
	for left <= right {
		mid := (left + right) / 2
		if mid*mid == x {
			return mid
		} else if mid*mid > x {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return right
}

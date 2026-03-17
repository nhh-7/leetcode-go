package times1

func mySqrt(x int) int {
	left, right := 0, x
	for left <= right {
		mid := (right-left)/2 + left
		if mid*mid > x {
			right = mid - 1
		} else if mid*mid < x {
			left = mid + 1
		} else {
			return mid
		}
	}
	return right
}

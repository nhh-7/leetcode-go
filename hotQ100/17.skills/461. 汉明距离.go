package skills

func hammingDistance(x int, y int) int {
	ans := 0
	for i := range 32 {
		a, b := (x>>i)&1, (y>>i)&1
		ans += a ^ b
	}
	return ans
}

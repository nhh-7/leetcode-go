package main

import (
	"fmt"
)

func test(t []int) {
	t[0] = 111
	fmt.Println(t, len(t), cap(t))
}

func main() {
	var s string = "1245"
	fmt.Println(s)

	sl := make([]int, 1, 11)
	fmt.Println(sl, len(sl), cap(sl))
	test(sl)
	fmt.Println(sl, len(sl), cap(sl))
}

package internal

func rand7() int {
	return 7
}

func rand10() int {
	for {
		row := rand7()
		col := rand7()
		x := (row-1)*7 + col
		if x <= 40 {
			return 1 + (x-1)%10
		}
	}
}

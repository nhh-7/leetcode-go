package main

import (
	"bufio"
	"fmt"
	"os"
)

func less(a, b []byte) bool {
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return a[i] < b[i]
		}
	}
	return false
}

func main() {
	in := bufio.NewReaderSize(os.Stdin, 1<<20)
	out := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer out.Flush()

	var n int
	var s string
	fmt.Fscan(in, &n)
	fmt.Fscan(in, &s)
	orig := []byte(s)
	best := make([]byte, n)
	copy(best, orig)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			for k := 1; i-k >= 0 && j+k < n; k++ {
				t := make([]byte, n)
				copy(t, orig)

				t[i], t[i-k] = t[i-k], t[i]
				t[j], t[j+k] = t[j+k], t[j]

				if less(t, best) {
					copy(best, t)
				}
			}
		}
	}
	fmt.Fprintln(out, string(best))
}

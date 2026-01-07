package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(100 * time.Millisecond)
	last := time.Now()

	for t := range ticker.C {
		fmt.Printf("间隔时间: %v\n", t.Sub(last))
		last = t
	}
}

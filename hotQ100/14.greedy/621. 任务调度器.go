package greedy

func leastInterval(tasks []byte, n int) int {
	maxCount := 0
	cnt := make(map[byte]int)
	for _, ch := range tasks {
		cnt[ch]++
		maxCount = max(maxCount, cnt[ch])
	}

	// 看最后一个桶中的任务数
	lastCount := 0
	for _, c := range cnt {
		if c == maxCount {
			lastCount++
		}
	}

	// 没满的时候
	notFullRes := (maxCount-1)*(n+1) + lastCount

	FullRes := len(tasks)
	return max(FullRes, notFullRes)
}

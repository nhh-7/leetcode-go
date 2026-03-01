package main

import (
	"context"
	"fmt"
	"strings"
	"sync"
	"time"
)

var (
	once sync.Once
	data string
)

func initialize() {
	fmt.Println("--- 执行初始化逻辑 ---")
	data = "Hello, GopherAI!"
}

func task(ctx context.Context) {
	fmt.Println(ctx.Deadline())
	select {
	case <-time.After(5 * time.Second):
		fmt.Println("5s")
	case <-ctx.Done():
		fmt.Println("context 倒计时到")
	}
}

func main() {
	var b strings.Builder
	b.WriteString("aa")
	b.WriteString(". ")
	fmt.Println(b.String())

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			once.Do(initialize)
			fmt.Printf("协程 %d 读取数据: %s\n", id, data)
		}(i)
	}
	wg.Wait()

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	go task(ctx)

	time.Sleep(3 * time.Second)
}

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

func test(t ...int) {
	for _, v := range t {
		fmt.Println(v)
	}
}

type node struct {
	left, right *node
	val         int
}

func createTree(nums []int) *node {
	if len(nums) == 0 {
		return nil
	}
	mid := len(nums) / 2
	root := &node{val: nums[mid]}
	root.left = createTree(nums[:mid])
	root.right = createTree(nums[mid+1:])
	return root
}

type listnode struct {
	next *listnode
	val  int
}

func createList(nums []int) *listnode {
	dummy := &listnode{}
	cur := dummy

	for _, v := range nums {
		cur.next = &listnode{val: v}
		cur = cur.next
	}
	return dummy.next
}

func partition(nums []int, left, right int) int {
	pivot := nums[left]
	i, j := left+1, right
	for {
		for i <= j && nums[i] < pivot {
			i++
		}
		for i <= j && nums[j] > pivot {
			j--
		}
		if i >= j {
			break
		}
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
	nums[left], nums[j] = nums[j], nums[left]
	return j
}

func main() {
	var b strings.Builder
	b.WriteString("aa")
	b.WriteString(". ")
	fmt.Println(b.String())

	test(1, 2, 3, 4, 5, 7)

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

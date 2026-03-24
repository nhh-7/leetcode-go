package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main1() {
	// 开启100个协程，顺序打印1-1000，且保证协程号1的，打印尾数为1的数字
	s := make(chan struct{})
	m := make(map[int]chan int, 100)
	for i := 1; i <= 100; i++ {
		m[i] = make(chan int)
	}

	for i := 1; i <= 100; i++ {
		go func(id int) {
			for {
				v := <-m[id]
				fmt.Println(v)
				s <- struct{}{}
			}
		}(i)
	}

	for i := 1; i <= 1000; i++ {
		id := i % 100
		if id == 0 {
			id = 100
		}
		m[id] <- i
		<-s
	}

	time.Sleep(10 * time.Second)
}

func main2() {
	// 三个goroutinue交替打印abc 10次
	ch1, ch2, ch3 := make(chan struct{}), make(chan struct{}), make(chan struct{})
	var wg sync.WaitGroup
	wg.Add(3)
	go func() {
		for i := 0; i < 10; i++ {
			<-ch1
			fmt.Println("a")
			ch2 <- struct{}{}
		}
		<-ch1
		wg.Done()
	}()

	go func() {
		for i := 0; i < 10; i++ {
			<-ch2
			fmt.Println("b")
			ch3 <- struct{}{}
		}
		wg.Done()
	}()

	go func() {
		for i := 0; i < 10; i++ {
			<-ch3
			fmt.Println("c")
			ch1 <- struct{}{}
		}
		wg.Done()
	}()

	ch1 <- struct{}{}
	wg.Wait()
	close(ch1)
	close(ch2)
	close(ch3)
	fmt.Println("end")
}

func main3() {
	// 交替打印奇偶数
	ch := make(chan struct{})
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		for i := 0; i < 10; i++ {
			ch <- struct{}{}
			if i%2 == 0 {
				fmt.Println("Even", i)
			}
		}
		wg.Done()
	}()

	go func() {
		for i := 0; i < 10; i++ {
			<-ch
			if i%2 == 1 {
				fmt.Println("Odd", i)
			}
		}
		wg.Done()
	}()
	wg.Wait()
}

func main4() {
	// 用单个channel实现0,1的交替打印
	ch := make(chan struct{})

	go func() {
		for {
			<-ch
			fmt.Println("0")
			ch <- struct{}{}
		}
	}()

	go func() {
		for {
			<-ch
			fmt.Println("1")
			ch <- struct{}{}
		}
	}()

	ch <- struct{}{}
	time.Sleep(2 * time.Second)
}

func main5() {
	// 使用两个Goroutine，向标准输出中按顺序按顺序交替打出字母与数字，输出是a1b2c3
	numCh := make(chan struct{})
	strCh := make(chan struct{})
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		for i := 'a'; i <= 'z'; i++ {
			fmt.Println(string(i))
			numCh <- struct{}{}
			<-strCh
		}
		wg.Done()
	}()

	go func() {
		for i := 1; i <= 26; i++ {
			<-numCh
			fmt.Println(i)
			strCh <- struct{}{}
		}
		wg.Done()
	}()

	wg.Wait()
	fmt.Println("finished")
}

func main6() {
	// 使用go实现1000个并发控制并设置执行超时时间1秒
	tasks := make(chan int, 1000)
	// 定义 ctx
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	var wg sync.WaitGroup

	// 启动 1000 个协程
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		tasks <- i
		go func(id int) {
			defer wg.Done()
			select {
			case <-ctx.Done():
				return
			default:
				fmt.Printf("goroutine id: %d\n", id)
			}
		}(i)
	}

	<-ctx.Done()
	fmt.Println("exec done")
	close(tasks)
	wg.Wait()
	fmt.Println("finish")
}

func main() {
	// 编写一个程序限制10个goroutine执行，每执行完一个goroutine就放一个新的goroutine进来
	var wg sync.WaitGroup
	ch := make(chan struct{}, 10)
	for i := 0; i < 20; i++ {
		wg.Add(1)
		ch <- struct{}{}
		go func(id int) {
			defer wg.Done()
			fmt.Printf("id: %d\n", id)
			<-ch
		}(i)
	}
	wg.Wait()
}

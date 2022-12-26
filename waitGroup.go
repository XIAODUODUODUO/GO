package main

import (
	"fmt"
	"sync"
	"time"
)

func wordker(id int) {
	fmt.Printf("worder %d starting\n", id)

	time.Sleep(time.Second)
	fmt.Printf("worder %d done\n", id)
}

func main() {
	//WaitGroup 用于等待这里启动的所有协程完成,如果 WaitGroup 显式传递到函数中，则应使用 指针
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1) // Add 添加需要等待的个数

		i := i
		//将 worker 调用包装在一个闭包中，可以确保通知 WaitGroup 此工作线程已完成。 这样，worker 线程本身就不必知道其执行中涉及的并发原语。
		go func() {
			defer wg.Done() //Done 来通知 WaitGroup 任务已完成
			wordker(i)
		}()
	}
	wg.Wait() // Wait 来等待所有 goroutine 结束
}

package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := 0; i <= 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	f("direct")

	// 使用 go f(s) 在一个协程中调用这个函数。 这个新的 Go 协程将会 并发地 执行这个函数。
	go f("goroutine")

	//为匿名函数启动一个协程
	go func(msg string) {
		fmt.Println("msg:" + msg)
	}("going")
	time.Sleep(time.Second)
	fmt.Println("done")
}

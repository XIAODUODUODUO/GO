package main

import (
	"fmt"
	"time"
)

func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Print("done")

	done <- true
}

func main() {

	done := make(chan bool, 1)
	//运行一个 worker 协程，并给予用于通知的通道。
	//程序将一直阻塞，直至收到 worker 使用通道发送的通知。
	go worker(done)

	<-done
}

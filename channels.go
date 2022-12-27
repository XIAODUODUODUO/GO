package main

import "fmt"

func main() {
	message := make(chan string) //make(chan val-type) 创建一个新的通道。 通道类型就是他们需要传递值的类型。

	//使用 channel <- 语法 发送 一个新的值到通道中。
	//在一个新的协程中发送 "ping" 到上面创建的 messages 通道中。
	go func() { message <- "ping" }()

	//使用 <-channel 语法从通道中 接收 一个值。
	//收到在上面发送的 "ping" 消息并将其打印出来。
	msg := <-message
	fmt.Println(msg)
}

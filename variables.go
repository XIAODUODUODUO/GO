package main

import "fmt"

func main() {

	var a = "initial" //var 声明 1 个或者多个变量。
	fmt.Println(a)

	var b, c int = 1, 2 //可以一次性声明多个变量。
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	var e int //声明后却没有给出对应的初始值时，变量将会初始化为 零值 。
	fmt.Println(e)

	f := "short" //:= 语法是声明并初始化变量的简写， 例如 var f string = "short"
	fmt.Println(f)
}

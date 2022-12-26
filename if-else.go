package main

import "fmt"

// Go 没有三目运算符， 即使是基本的条件判断，依然需要使用完整的 if 语句。
func main() {

	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}
	//在条件语句之前可以有一个声明语句；在这里声明的变量可以在这个语句所有的条件分支中使用。
	//注意，在 Go 中，条件语句的圆括号不是必需的，但是花括号是必需的。
	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
}

package main

import (
	"fmt"
	"math"
)

const s string = "constant" //const 用于声明一个常量。const 语句可以出现在任何 var 语句可以出现的地方

func main() {
	fmt.Println(s)

	//常数表达式可以执行任意精度的运算
	const n = 500000000

	//数值型常量没有确定的类型，直到被给定某个类型，比如显式类型转化。
	//一个数字可以根据上下文的需要（比如变量赋值、函数调用）自动确定类型。 举个例子，这里的 math.Sin 函数需要一个 float64 的参数，n 会自动确定类型。
	const d = 3e20 / n
	fmt.Println(d)

	fmt.Println(int64(d))

	fmt.Println(math.Sin(n))
}

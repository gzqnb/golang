package main

import "fmt"

const pi = 3.1415

const (
	statusOK = 200
	notTFOUNT = 404 
)

const (
	n1 = 100
	// 没有写值默认和上一行一样
	n2
	n3
)
// iota 计数器 类似枚举
const(
	a1 = iota //0
	a2 = iota //1 每新增一行加1
	a3
)

const(
	b1 = iota // 当const出现时第一次iota
	b2
	_   // 匿名变量
	b3
)

func main()  {
	fmt.Println(a1)
	fmt.Println(a2)
	fmt.Println(a3)
	fmt.Println(b1)
	fmt.Println(b2)
	fmt.Println(b3)
}
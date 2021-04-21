package main

import "fmt"
// var name string
// var age int
// var isOk bool
// 批量声明
var (
	name string  //""
	age int     //0
	isOk bool //false
)
func main() {
    name = "理想"
	 age = 16
	 isOk = true
	 fmt.Printf("name:%s",name) //%s占位符
	 fmt.Println(age)
	 fmt.Print(isOk)
     // var代表声明变量 一个作用域只能声明一次变量
	 var s1  = "aaa"
	 fmt.Println(s1)
	 //简短声明只能在函数里用
	 s2 := "aaa"
	 fmt.Print(s2)
	 

}
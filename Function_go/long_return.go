package main

import (
	"fmt"
)

func test() (int, string, bool) {
	return 666, "hello world", true
}

//go语言推荐写法
func test2() (a, b, c int) {
	a, b, c = 66, 88, 99
	return
}

func main() {
	a, b, c := test()
	fmt.Printf("a=%d,b=%s,c=%t", a, b, c)
	aa, bb, cc := test2()
	fmt.Printf("\naa=%d,bb=%d,cc=%d", aa, bb, cc)
}

/*
总结：Go语言函数定义说明：
1) func: 函数由关键字func开始声明
2) FuncName: 函数名称，根据规定，函数名首字母大写为pubilc,小写为private
3) 参数列表： 函数可以有0个或者多个参数，参数格式为： 变量名 类型，如果有多个参数通过
	逗号分隔，不支持默认参数
4) 返回类型：
	1) 上面返回值声明了两个变量名(命名返回参数)，这个不是必须，可以只有类型没有变量名
	2) 如果只有一个返回值且不声明返回值变量，那么你可以省略，包括返回值的括号：func test() int{
	3) 如果没有返回值，那么就直接省略最后的返回信息
	4) 如果有返回值，那么必须在函数的内部添加return语句






*/

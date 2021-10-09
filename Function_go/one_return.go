package main

import (
	"fmt"
)

//第一种带一个返回值的函数
func retu_fun() int {
	return 666
}

//带返回值的函数的常规形式
func retur_fun2() (result int) {
	result = 888
	return
}

func main() {
	//接收函数返回值的方法：
	//1.
	var a int
	a = retu_fun()
	fmt.Println(a)

	//2.
	b := retu_fun()
	fmt.Println(b)

	c := retur_fun2()
	fmt.Println(c)

}

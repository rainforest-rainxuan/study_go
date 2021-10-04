package main

import "fmt"

type MyInt64 int64

func test1() {
	var a int = 1
	var b int64
	//b = a
	var c MyInt64
	//c = a	 //cannot use a (type int) as type MyInt64 in assignment
	c = MyInt64(a)
	b = int64(a)
	fmt.Println(b)
	fmt.Println(c)
}

func main() {
	test1()
}

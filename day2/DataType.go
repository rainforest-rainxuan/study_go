package main

import "fmt"

func test1() {
	var a int = 1
	var b int64
	//b = a
	b = int64(a)
	fmt.Println(b)
}

func main() {
	test1()
}

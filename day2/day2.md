# Day2
## go语言不支持隐式转换，类型别名都不行

```go
package main

import "fmt"

func test1() {
	var a int = 1
	var b int64
	//b = a  err:cannot use a (type int) as type int64 in assignment
    b = int64(a)
	fmt.Println(b)
}

func main() {
	test1()
}
```

## 有指针，但不支持指针的运算

## string字符串zero值为空串，不是nil


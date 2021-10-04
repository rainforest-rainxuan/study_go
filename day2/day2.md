# Day2
## go语言不支持隐式转换，类型别名都不行

```go
package main

import "fmt"

type MyInt64 int64

func test1() {
	var a int = 1
	var b int64
	//b = a  //err:cannot use a (type int) as type int64 in assignment
    b = int64(a)
    var c MyInt64;
    //c = a;  //cannot use a (type int) as type MyInt64 in assignment
    c = MyInt64(a)
	fmt.Println(b)
    fmt.Println(c)
}

func main() {
	test1()
}
```

## 有指针，但不支持指针的运算

## string字符串zero值为空串，不是nil


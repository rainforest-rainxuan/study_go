package main

import "fmt"

func main() {
	citys := [7]string{"北京", "上海", "广东", "深圳", "南京", "成都", "武汉"}
	//基于citys创建一个新数组
	names := [3]string{}
	names[0] = citys[0]
	names[1] = citys[1]
	names[2] = citys[2]
	fmt.Println("names=", names)

	//切片可以根据一个数组创建一个新的数组
	names1 := citys[0:3] //左闭右开
	fmt.Println("names1=", names1)

	//基于一个数组的切片，应该是对这个数组进行了浅拷贝，如果切片数组的值改变，那么原来的数组也会改变
	names1[2] = "hello"
	fmt.Println("改变了切片数组names1:", names1)
	fmt.Println("改变了切片数组的原始数组citys:", citys)

	//如果想要切片得到的数组完全独立于原数组，深拷贝，可以靠copy()函数实现

	//如果从0开始截取，那么冒号左边的0可以省略
	names3 := citys[:5]
	fmt.Println("names3=", names3)

	//从中间截取
	names4 := citys[3:5] //左边右开，3  4  "深圳"  "南京"
	fmt.Println("names4=", names4)

	//如果是从数组的某个点截取到数组的末尾，那么冒号右边可以省略
	names5 := citys[3:]
	fmt.Println("names5=", names5)

	//如果想得到数组全部的切片，不过我认为这就没什么意义了
	names6 := citys[:]
	fmt.Println("names6=", names6)

	//截取数组某个元素 等价于 array[index] 同样的，我也认为这没什么意义，不如直接写array[index]
	names7 := citys[5:6]
	fmt.Println("names7=", names7)

	//字符串形成的切片
	names8 := "hello world!"[3:7]
	fmt.Println("names8=", names8)

	//可以根据make，创建一个空的，指定大小的切片
	sub := make([]string, 0, 20) //第一个参数是type 第二个参数是当前的大小，第三个参数是切片的容量大小,非必须，如果没有传，默认第二个参数一样
	fmt.Println("len(sub)=", len(sub), "cap(sub)=", cap(sub))
	//sub[0] = "hello"
	sub = append(sub, "hello")
	fmt.Println("sub=", sub)
	fmt.Println("len(sub)=", len(sub), "cap(sub)=", cap(sub))

	//copy()深拷贝
	str := make([]string, len(citys))
	copy(str, citys[:])
	str[1] = "台湾"
	fmt.Println("str=", str)
	fmt.Println("citys=", citys)
}

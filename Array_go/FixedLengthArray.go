package main

import (
	"fmt"
)

func main() {
	arr := [6]int{1, 2, 3, 4, 5, 6} //初始化列表方式给数组赋值
	for i := 0; i < len(arr); i++ {
		fmt.Printf("arr[%d]=%d\t", i, arr[i])
	}
	fmt.Printf("\n")
	arr[3] = 88
	for i, data := range arr {
		fmt.Printf("arr[%d]=%d\t", i, data)
	}
	fmt.Printf("\n")

	//循环初始化
	var arr2 [5]int
	for i := 0; i < 5; i++ {
		arr2[i] = i
	}
	for i, data := range arr2 { //data:只是arr2的副本，和arr2无关
		fmt.Printf("arr2[%d]=%d\t", i, data)
	}
	fmt.Printf("\n")

	//go语言居然支持相同数组的赋值
	arr3 := arr2

	for i, data := range arr3 {
		fmt.Printf("arr3[%d]=%d\t", i, data)
	}
	fmt.Printf("\n")

	// err:cannot use arr (type [6]int) as type [7]int in assignment
	var arr5 [7]int = [7]int{11, 22, 33, 55, 66, 77, 88}
	// arr5 = arr
	// for i, data := range arr5 {
	// 	fmt.Printf("arr[%d]=%d\t", i, data)
	// }
	// fmt.Printf("\n")

	//arr = arr5	 //err

	//总结，不同大小的数组的不能相互赋值

	fmt.Println("arr[6]=", arr5[6])

	var arr6 [5]int
	arr6 = arr3
	fmt.Println("arr6[3]=", arr6[3])

}

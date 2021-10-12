package main

import "fmt"

func main() {
	citys := []string{"北京", "上海", "广东", "深圳", "南京", "成都", "武汉"}
	for i, data := range citys {
		fmt.Println("i=", i, "city:", data)
	}

	Citys := append(citys, "合肥")
	for _, data := range Citys {
		fmt.Println("Citys=", data)
	}

	for i, data := range citys {
		fmt.Println("i=", i, "city:", data)
	}

	//一个切片不止有长度len()，还有容量cap()
	fmt.Println("添加元素之前的长度：", len(citys), "添加元素之前的容量:", cap(citys))

	citys = append(citys, "大理")
	fmt.Println("赋值给自己一个新值后的citys=", citys)

	fmt.Println("添加元素后的长度:", len(citys), "添加元素后的容量:", cap(citys))

	citys = append(citys, "杭州")
	fmt.Println("再次添加一个元素后的切片:", citys)
	fmt.Println("再次添加一个元素的长度:", len(citys), "再次添加一个元素的容量:", cap(citys))

	//为了防止反复添加元素，切片会自动扩容，当所需空间的少于1k时会自动扩容为两倍，超过应该一次扩容1.5倍或者更低
	var nums []int
	for i := 0; i < 50; i++ {
		nums = append(nums, i)
		fmt.Println("len=", len(nums), "cap=", cap(nums))
	}
}

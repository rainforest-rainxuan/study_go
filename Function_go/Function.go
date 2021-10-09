package main

import (
	"fmt"
	_ "time" //匿名包，可以不用使用
)

//无参无返回值的函数
func void_fun() {
	a := 10
	fmt.Println("a= ", a)
}

//有参无返回值的函数
func paramter_fun(b int) {
	fmt.Println("b= ", b)
}

//多个形参无返回值的函数
func double_paramter_fun(a, b int, str string, ft float64) {
	fmt.Printf("c=%d,d=%d,str=%s,ft=%f\n", a, b, str, ft)
}

//建议每个形参都要注明类型，比较清晰，就是不简洁

//不定参数的形参  args ...type	可以有零个或者多个参数
func variable_param_fun(agrs ...int) {
	for i, data := range agrs {
		fmt.Printf("args[%d]=%d\n", i, data)
	}
	fmt.Println("+++++++++++++++++++==++++++++==+++++++")
	for i := 0; i < len(agrs); i++ {
		fmt.Printf("args[%d]=%d\n", i, agrs[i]) //切片方式
	}

	fmt.Printf("len=%d\n", len(agrs))
}

// 不定参数的形参只能出现在参数列表的最后一个，否则报下面的错
// 不定参数前面有形参必须传递实参
// func test(a ...int, b int) { //syntax error: cannot use ... with non-final parameter a

// }

func main() {
	//start := time.Now()
	void_fun()
	b := 666
	paramter_fun(b)
	//time.Sleep(time.Second)
	a, b, str, ft := 88, 66, "牛皮666", 3.14
	double_paramter_fun(a, b, str, ft)
	variable_param_fun(2, 3, 4, 5, 6)
	//end := time.Now()
	//_ = time.Now(); 匿名变量不会报错
	//total := end.Sub(start)
	//fmt.Println("total time is ", total)

	fmt.Println("----------------------")
	/*************/
	var_para_fun(1, 2, 3, 4, 5)
}

/*********************/
//不定参数的传递

func var_para_fun(args ...int) {
	//1. 全部参数的传递
	test(args...)

	fmt.Println("-------------")

	//2.部分参数的传递
	test(args[:3]...) //传递args[0]到args[2]

	fmt.Println("-------------")
	test(args[3:]...) //传递args[3]包括args[3]到最后的值
}

func test(tmp ...int) {
	for _, data := range tmp {
		fmt.Printf("data=%d\n", data)
	}
}

/*********************/

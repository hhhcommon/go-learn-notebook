package main

import "fmt"

func main() {
	var a int = 1
	var ptr *int = &a
	fmt.Println("a的值是：", a)
	fmt.Println("a 的指针ptr, ptr的值是：", ptr)
	fmt.Printf("ptr指向的值：%v\n", *ptr)

	fmt.Println("我要开始修改 *ptr 的值了：：：")

	*ptr = 999

	fmt.Println("修改后a的值是：", a)

	fmt.Println("我要做一些骚操作了：：：")

	// var b int = *ptr
	b := *ptr

	b = 20

	fmt.Println("修改后a的值是：", a)
	fmt.Println("b的值是：", b)

	fmt.Println(97%7, 97/7, 5.0/9*(1000-100))

}

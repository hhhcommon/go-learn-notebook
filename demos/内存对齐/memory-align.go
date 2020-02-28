package main

import (
	"fmt"
	"unsafe"
)

type Panda struct {
	a int32
	b int64
	c bool
	d string
	e float32
	f float64
}

func PrintEveryTypeMoemorySize()  {
	fmt.Printf("int32 的内存占用为：%d, 分配系数为：%d\n", unsafe.Sizeof(int32(0)), unsafe.Alignof(int32(0)))
	fmt.Printf("int64 的内存占用为：%d, 分配系数为：%d\n", unsafe.Sizeof(int64(0)), unsafe.Alignof(int64(0)))
	fmt.Printf("bool 的内存占用为：%d, 分配系数为：%d\n", unsafe.Sizeof(bool(false)), unsafe.Alignof(bool(false)))
	fmt.Printf("string 的内存占用为：%d, 分配系数为：%d\n", unsafe.Sizeof(string("")), unsafe.Alignof(string("")))
	fmt.Printf("float32 的内存占用为：%d, 分配系数为：%d\n", unsafe.Sizeof(float32(0)), unsafe.Alignof(float32(0)))
	fmt.Printf("float64 的内存占用为：%d, 分配系数为：%d\n", unsafe.Sizeof(float64(0)), unsafe.Alignof(float64(0)))
	
}

func PrintStructMemorySize()  {
	fmt.Printf("Panda结构体的内存占用为：%d, 分配系数为：%d\n", unsafe.Sizeof(Panda{}), unsafe.Alignof(Panda{}))
}

func main()  {
	PrintStructMemorySize()

	PrintEveryTypeMoemorySize()

	
}
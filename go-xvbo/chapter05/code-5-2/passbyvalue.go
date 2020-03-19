// 5-2：参数值传递

package main

import "fmt"

// Data 用于测试值传递效果的结构体
type Data struct {
	complex  []int      // 测试切片在参数传递中的效果
	instance InnerData  // 实例分配的 innerData
	ptr      *InnerData // 将 ptr 声明为 InnerData 的指针类型
}

// InnerData 代表各种结构字段
type InnerData struct {
	a int
}

// passByValue 值传递测试函数
func passByValue(inFunc Data) Data {

	fmt.Printf("inFunc value: %+v\n", inFunc)

	fmt.Printf("inFunc ptr: %p\n", &inFunc)

	// 返回的过程将发生复制
	return inFunc
}

func passPtr(inPtrFunc *Data) *Data {
	data := *inPtrFunc
	fmt.Printf("inPtrFunc value: %+v\n", data)
	fmt.Printf("inPtrFunc ptr: %p\n", &data)
	return &data
}

func main() {
	in := Data{
		complex: []int{1, 2, 3},
		instance: InnerData{
			5,
		},
		ptr: &InnerData{1},
	}

	fmt.Printf("in value: %+v\n", in)

	fmt.Printf("in ptr: %p\n", &in)

	// out := passByValue(in)
	outPtr := passPtr(&in)
	out := *outPtr
	fmt.Printf("out value: %+v\n", out)
	fmt.Printf("out ptr: %p\n", &out)

}

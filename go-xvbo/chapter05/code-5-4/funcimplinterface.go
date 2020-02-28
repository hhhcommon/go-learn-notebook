// 函数实现接口
package main

import "fmt"

// Struct 结构体
type Struct struct {
}

// Invoker 调用器接口
type Invoker interface {
	Call(interface{})
	// Hello(interface{})
}

// Call Struct实现 Invoker 调用器的接口中的Call方法。
func (s *Struct) Call(p interface{}) {
	fmt.Println("from struct", p)
}

// 也就是说，Struct 实现 Invoker接口只需要声明一个 针对 Struct接口体的Call 方法就行了。
func main() {
	var invoker Invoker

	s := new(Struct)

	// 这里 只有Struct实现了 Invoker接口的所有成员 ，Struct 和 Invoker才是相同的类型。
	// 例如 在 接口里增加一个Hello成员，Struct 不去实现 Hello方法，就不能赋值给 Invoker类型的变量。
	invoker = s
	s.Call("World !")
	invoker.Call("Hello")
}

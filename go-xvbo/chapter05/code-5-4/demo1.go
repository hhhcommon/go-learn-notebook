// 函数体实现接口
package main

import "fmt"

type Invoker interface {
	Call(interface{})
}

// FuncCaller 函数定义为类型
type FuncCaller func(interface{})

// Call 实现Invoker的Call
// 仅仅一个 Call 名字和参数一样就算是 让 FuncCaller类型的函数实现了 Invoker接口？
// 骚，真的骚
func (f FuncCaller) Call(p interface{}) {
	f(p)
}

// 一下是HTTP包里的例子 让函数类型实现接口
// 保持 函数名和函数入参类型数量保持一致即可
// type Handler interface {
// 	ServeHTP(ResponseWriter, *Request)
// }

// type HandlerFunc func(ResponseWriter, *Request)

// func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request)  {
// 	f(w,r)
// }

func main() {
	var invoker Invoker

	invoker = FuncCaller(func(v interface{}) {
		fmt.Println("from Fucntion", v)
	})

	invoker.Call("hello")


	str := "Hello World!"

	foo := func(s string) {
		str = "hello dude!"
	}

	foo(str)

	fmt.Println(str)
}

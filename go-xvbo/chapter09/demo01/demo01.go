package main

import (
	"fmt"
	"time"
)

func running() {
	var times int
	// 构建一个无限循环
	for {
		times++
		fmt.Println("tick", times)
		// 延时一秒
		time.Sleep(time.Second)
	}
}

func main() {
	// 开启一个 goroutine 运行 running 函数
	go running()

	// 同时在 main 函数这个本身就在默认的 goroutine 里运行的函数里
	// 等待用户输入一个字符。
	// 只要用户输入任意字符，main 函数所在的 goroutine 停止运行
	// 于此同时，在这个 goroutine 里面创建的 goroutine 也会随之停止。
	var input string
	fmt.Scanln(&input)
}

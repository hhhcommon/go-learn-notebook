package main

import (
	"strings"
	"fmt"
)

// StringProccess 字符串的链式处理
func StringProccess(list []string, chain []func(string) string) {
	// 遍历每一个字符串
	for index, str := range list {
		// 每一个需要处理的字符串
		result := str
		// 遍历每一个处理链
		for _, proc := range chain {
			// 执行处理得到结果
			result = proc(result)
		}
		// 将结果重新赋值回原来的 List 里面
		list[index] = result
	}
}

// RemovePrefix 自定义的移除浅醉的处理函数
func RemovePrefix(str string) string {
	return strings.TrimPrefix(str, "go")
}

func main() {
	list := []string{
		"go scanner",
		"go parser",
		"go compiler",
		"go printer",
		"go formatter",
	}

	chain := []func(string) string{
		RemovePrefix,
		strings.TrimSpace,
		strings.ToUpper,
	}

	StringProccess(list, chain)

	for _, str := range list {
		fmt.Println(str)
	}
}

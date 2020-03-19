// 5-1：将秒解析为时间单位

package main

import "fmt"

const (
	SecondsPerMinute = 60

	SecondsPerHour = SecondsPerMinute * 60

	SecondsPerDay = SecondsPerHour * 24
)

func resolveTime(seconds int) (day int, hour int, minut int) {
	day = seconds / SecondsPerDay
	hour = seconds / SecondsPerHour
	minut = seconds / SecondsPerMinute

	return
}

func main() {
	// 将返回值作为打印函数的参数
	fmt.Println(resolveTime((1000)))
	// 只获取小时和分钟
	_, hour, minute := resolveTime(18000)
	fmt.Println(hour, minute)
	// 只获取天
	day, _, _ := resolveTime(90000)
	fmt.Println(day)
}

// output：
// 0 0 16
// 5 300
// 1

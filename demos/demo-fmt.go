package main

import (
	"fmt"
	"math"
)

// 通用输出方法：
// %v 输出 值
// %#v go语法样式的 值
// %T  go语法类型的 值
// %%  百分号的字面量， 不代表任何值
func main() {
	fmt.Println("int==============")

	integer := 23

	fmt.Println(integer)

	fmt.Printf("%v\n", integer)

	fmt.Printf("%#v\n", integer)

	fmt.Printf("%d\n", integer)

	fmt.Printf("%T\n", integer)

	fmt.Printf("%T %T\n", integer, &integer)

	fmt.Println("boolean==============")
	// %t 专门字符串形式输出boolean值，输出 true 或 false
	truth := true

	fmt.Printf("%v %t\n", truth, truth)

	fmt.Println("int用不同的进位制输出==============")

	answer := 42

	fmt.Printf("值：%v 十进制：%d 16进制：%x 8进制：%o 二进制：%b\n", answer, answer, answer, answer, answer)

	pi := math.Pi

	fmt.Printf("值：%v  大指数否者使用%%f ： %g  两位小数：%0.2f 4位整数两位小数共六位： (%6.2f) 科学计数法：%e\n", pi, pi, pi, pi, pi)

	fmt.Printf("%T\n", pi)

	fmt.Println("虚数==============")

	point := 110.7 + 22.5i

	fmt.Printf("%v %g %.2f %.2e\n", point, point, point, point)

	fmt.Println("Unicode==============")
	smile := '🙂'
	fmt.Printf("值：%v  十进制： %d  字符：%c  转义字符：%q  十六进制Unicode编码字符：%U  十六进制Unicode编码字符＋可打印字符：%#U\n", smile, smile, smile, smile, smile, smile)

	fmt.Println("原格式字符串==============")
	placeholders := `foo "bar"`

	fmt.Printf("值：%v  as-is：%s  转义字符：%q  块字符：%#q\n", placeholders, placeholders, placeholders, placeholders)

	fmt.Println("字典格式化输出值==============")
	isLegume := map[string]bool{
		"peanut":    true,
		"dachshund": false,
	}

	fmt.Printf("默认输出值：%v  使用go源码格式化后的map值：%#v\n", isLegume, isLegume)

	fmt.Println("结构体格式化输出==============")
	person := struct {
		Name string
		Age  int
	}{"Kim", 22}

	fmt.Printf("默认输出值：%v 使用name字段排序%+v 使用go源码格式化后的struct值：%#v\n", person, person, person)
	fmt.Println("指针格式化输出==============")
	pointer := &person

	fmt.Printf("值：%v 将指针输出成16进制：%p\n", pointer, (*int)(nil))
	fmt.Printf("值：%v 将指针输出成16进制：%p\n", pointer, pointer)
	fmt.Println("切片格式化输出==============")
	fmt.Println("数组格式化输出==============")
	fmt.Println("字节数组格式化输出==============")
}

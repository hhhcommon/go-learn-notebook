// 稀疏数组

// 有些二维数组型数据结构里面，有很多是零填充的空位。
// 为了压缩数据，我们使用三元组压缩数据
// 1. 二维数组的行列 和默认值
// 2. 记录 其他值 和 这个值所在的行列。

// 稀疏数组系列化

// 序列化的相关东西主要涉及文件读写。
// 这部分我们接下来进行目前着眼数据结构。
package main

import (
	"fmt"
)

// Node 定义压缩字段使用的三元组
type Node struct {
	row   int
	colum int
	value int
}

func showArr(arr *[3][3]int) {
	// 拿出每一行
	for _, row := range arr {
		// 打印一行里面的每一个值
		for _, value := range row {
			fmt.Printf("%d  ", value)
		}
		fmt.Println()
	}
}

func main() {

	// 1. 创建一个原始数组代表棋盘
	var arr [3][3]int // 这里棋盘会被0初识化填充，0表示没有棋子

	arr[0][0] = 1 // 数字1代表白子
	arr[1][2] = 2 // 数字2代表黑子

	fmt.Println("看一下原来的二维数组：")
	showArr(&arr)

	// 转换成稀疏数组
	// 2. 创建一个又稀疏节点组成的数组
	var sparseArr []Node
	// 稀疏数组的第一个元素是记录整个二维数据的宽高和默认值
	node := Node{
		row:   3,
		colum: 3,
		value: 0,
	}
	sparseArr = append(sparseArr, node)

	// 3. 转换：接下来遍历二维数据将特数据找出来放进稀疏数组
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			if arr[i][j] != 0 {
				sparseArr = append(sparseArr, Node{
					row:   i,
					colum: j,
					value: arr[i][j],
				})
			}
		}
	}

	// 4. 测试：打印看看那数据对不对
	fmt.Println("接下来看看转换成稀疏数组后 记录了哪些值。")
	for _, n := range sparseArr {
		fmt.Printf("row: %d, colum: %d, value: %d\n", n.row, n.colum, n.value)
	}
}

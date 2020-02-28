package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

// main 客户端
// 客户端可以发送单行数据然后退出
// os。Stdin是落到一个文件上的，代表标准输入【终端】
func main() {
	// 创建一个连接
	conn, err := net.Dial("tcp", "0.0.0.0:8888")
	if err != nil {
		fmt.Println("Client dial err =", err)
		return
	}
	// 创建一个标准输入
	reader := bufio.NewReader(os.Stdin)

	// fmt.Printf("%v\n", conn)
	// 从终端按行读取输入，并准备发送给服务端。
	line, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("readString err =", err)
	}

	// 将line 发送给 服务器
	n, err := conn.Write([]byte(line))

	if err != nil {
		fmt.Println("conn.Write err =", err)
		return
	}

	fmt.Println("n = ", n)

}

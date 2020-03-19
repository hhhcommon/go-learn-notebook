package main

import (
	"fmt"
	"net"
)

func main() {
	listen, err := net.Listen("tcp", "0.0.0.0:8888")

	if err != nil {
		fmt.Println("listen err=", err)
	}

	defer listen.Close()

	for {
		fmt.Println("Wating for connection...")

		conn, err := listen.Accept()

		if err != nil {
			fmt.Println("Accept() err=", err)
		} else {
			fmt.Println("Accept() success con = %v\n Client IP = %v\n", conn, conn.RemoteAddr().String())
		}
	}
}

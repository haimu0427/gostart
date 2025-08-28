package main

import (
	"fmt"
	"net"
)

func process(conn net.Conn) {
	// 循环接受客户端数据
	defer conn.Close()
	for {
		//创建切片
		buf := make([]byte, 1024)
		fmt.Println("Waiting for client" + conn.RemoteAddr().String() + " to send data...")
		n, err := conn.Read(buf)
		if err != nil {

			if err.Error() == "EOF" {
				fmt.Println("Client disconnected:", conn.RemoteAddr().String())
				return
			} else {
				fmt.Println("Error reading data:", err)
				return
			}
		}
		fmt.Print(string(buf[:n]))

	}
}

func main() {
	fmt.Println("server start...")
	listener, err := net.Listen("tcp", "0.0.0.0:8888")
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
	fmt.Printf("Listener: %v\n", listener)
	defer listener.Close()

	//等待客户端连接
	for {
		fmt.Println("Waiting for client connection...")
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
		} else {
			fmt.Printf("Client connected: %v\n", conn.RemoteAddr().String())
			go process(conn) // 启动一个goroutine处理连接
		}

	}

}

package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	// 连接到服务器
	conn, err := net.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		return
	}
	defer conn.Close()
	fmt.Printf("Client connected to server: %v\n", conn.RemoteAddr().String())
	fmt.Println("Client local address:", conn.LocalAddr().String())
	// 读取用户输入并发送到服务器
	reader := bufio.NewReader(os.Stdin)
	line, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}
	exit := strings.Trim(line, "\r\n")
	if exit == "exit" {
		fmt.Println("Client exiting...")
		return
	}

	n, err := conn.Write([]byte(line))
	if err != nil {
		fmt.Println("Error sending data:", err)
		return
	}
	fmt.Printf("Sent %d bytes to server\n", n)
}

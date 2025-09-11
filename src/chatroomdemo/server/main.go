package main

import (
	"chatroom/common/message"
	"encoding/binary"
	"fmt"
	"net"
)

func readPkg(conn net.Conn) (mes message.Message, err error) {
	buf := make([]byte, 512)
	fmt.Println("服务器正在等待客户端发送信息...")
	n, err := conn.Read(buf[0:4]) //阻塞等待客户端发送信息
	if err != nil {
		fmt.Println("conn.Read err=", err)
		return
	}
	fmt.Println("接收到客户端发送的信息：", string(buf[:n]), len(buf[:n]))
	var pkgLen uint32
	binary.BigEndian.PutUint32(buf[0:4], pkgLen)

	//根据pkgLen读取消息内容

	return mes, nil
}

func process(conn net.Conn) { //套接字是一种引用类型
	//这里需要延时关闭conn
	defer conn.Close()

	//循环读取客户端发送的信息
	for {
		//读取数据包, 封装为一个函数

	}

}
func main() {
	fmt.Println("服务器在8889端口监听...")
	lister, err := net.Listen("tcp", "0.0.0.0:8889")
	if err != nil {
		fmt.Println("net.Listen err=", err)
		return
	}
	defer lister.Close()
	for {
		fmt.Println("等待客户端来连接...")
		conn, err := lister.Accept()
		if err != nil {
			fmt.Println("lister.Accept err=", err)
		}

		//一旦连接成功，则启动一个协程和客户端保持通讯
		go process(conn)
	}

}

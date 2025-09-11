package main

import (
	"chatroom/common/message"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
	_ "time"
)

func login(userId int, userPwd string) (err error) {

	// //需要定下协议...

	// // Simulate a login request
	// fmt.Println("Logging in:", userId)
	// time.Sleep(1 * time.Second) // Simulate network delay

	// // Return a mock token
	// fmt.Println("Login successful!")

	// 1.  link to server
	conn, err := net.Dial("tcp", "127.0.0.1:8889")
	if err != nil {
		fmt.Println("net.Dial err=", err)
		return
	}
	defer conn.Close()
	// 2. send message to server by conn
	var mes message.Message
	mes.Type = message.LoginMesType
	// 3. create LoginMes
	var loginMes message.LoginMes
	loginMes.UserId = userId
	loginMes.UserPwd = userPwd
	//4. serialize loginMes
	data, err := json.Marshal(loginMes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}
	//5. give data to mes.Data
	mes.Data = string(data)
	//6. serialize mes
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}
	//7. send data to server
	// 7.1 send the length of data
	var pkgLen = uint32(len(data))
	buf := make([]byte, 4)
	binary.BigEndian.PutUint32(buf[0:4], pkgLen)
	n, err := conn.Write(buf[0:4])
	if err != nil || n != 4 {
		fmt.Println("conn.Write err=", err)
		return
	}
	// 7.2 send the actual data
	_, err = conn.Write(data)
	if err != nil {
		fmt.Println("conn.Write err=", err)
		return
	}
	fmt.Println("客户端发送登录消息成功...")
	fmt.Println("len(data)=", len(data))
	return nil
}

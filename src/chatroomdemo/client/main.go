package main

import (
	"fmt"
)

var userId int
var userPwd string

func main() {
	// 接受用户的选择
	var choice int
	// 判断是否继续显示菜单
	var loop = true
	for loop {
		fmt.Println("欢迎登录多人聊天系统")
		fmt.Printf("\t\t\t 1. 登陆聊天系统\n")
		fmt.Println("\t\t\t 2. 注册用户")
		fmt.Print("\t\t\t 3. 退出系统\n")
		fmt.Print("\t\t\t 请选择(1-3): ")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			// 处理登录
			fmt.Println("正在登录...")
			loop = false
		case 2:
			// 处理注册
			fmt.Println("正在注册...")
			loop = false
		case 3:
			// 处理退出
			fmt.Println("退出系统...")
			loop = false
			//or os.Exit(0)
		default:
			fmt.Println("无效选择，请重新输入.")
		}
	}

	if choice == 1 {
		fmt.Println("请输入用户名:")
		fmt.Scanln(&userId)
		fmt.Println("请输入密码:")
		fmt.Scanln(&userPwd)
		err := login(userId, userPwd)
		if err != nil {
			fmt.Println("登录失败:", err)
		} else {
			fmt.Println("登录成功!")
		}
	} else if choice == 2 {
		fmt.Println("感谢使用，再见!")
	}

}

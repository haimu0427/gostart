package main

import (
	"flag"
	"fmt"
)

func main() {
	var user string
	var pwd string
	var host string
	var port int
	//&user -u 后面的参数值
	//u  -u参数
	//"" 默认值
	//"username for login" 说明
	flag.StringVar(&user, "u", "", "username for login")
	flag.StringVar(&pwd, "pwd", "", "password for login")
	flag.StringVar(&host, "h", "localhost", "host for login")
	flag.IntVar(&port, "p", 3306, "port for login")
	flag.Parse()
	fmt.Printf("user: %s, pwd: %s, host: %s, port: %d\n", user, pwd, host, port)
}

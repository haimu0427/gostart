package message

//消息结构体
const (
	LoginMesType    = "LoginMes"    //登录消息
	LoginResMesType = "LoginResMes" //登录响应消息
)

type Message struct {
	Type string `json:"type"`
	Data string `json:"data"`
}

type LoginMes struct {
	UserId   int    `json:"userId"`   //用户名
	UserPwd  string `json:"userPwd"`  //用户密码
	UserName string `json:"userName"` //用户名
}

type LoginResMes struct {
	Code  int    `json:"code"`  //返回状态码 500表示用户未注册 200表示登录成功
	Error string `json:"error"` //返回错误信息,

}

type Type struct {
}

type Data struct {
}

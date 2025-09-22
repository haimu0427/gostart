package main

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type myClaims struct {
	Foo string `json:"foo"`
	jwt.RegisteredClaims
	// 使用匿名实现 Embedding RegisteredClaims to include standard claims
}

func main() {
	// 定义用于签名令牌的密钥字符串
	mySigningKey := []byte("fangguowoba")

	// 创建一个新的令牌对象，指定签名方法和声明
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, myClaims{ //也可以使用map, 但是不如结构体方便
		Foo: "bar",
		RegisteredClaims: jwt.RegisteredClaims{
			NotBefore: jwt.NewNumericDate(time.Now()),
			// 设置令牌在当前时间之后生效
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(10 * time.Minute)),
			Issuer:    "first try",
		},
	})
	fmt.Println("Token:", token)
	fmt.Println("")

	// 使用指定的密钥字符串签名令牌并获得完整的编码令牌
	tokenString, err := token.SignedString(mySigningKey)
	if err != nil {
		fmt.Println("Error signing token:", err)
		return
	}
	// 输出签名后的令牌字符串, 可用于客户端传输
	fmt.Println("Signed Token:", tokenString)
	fmt.Println("")

	// 解析令牌并验证签名
	token2, err := jwt.ParseWithClaims(tokenString, &myClaims{},
		func(token2 *jwt.Token) (interface{}, error) {
			return mySigningKey, nil
		})
	if err != nil {
		fmt.Println("Error parsing token:", err)
		return
	}
	// 验证令牌的有效性
	if claims, ok := token2.Claims.(*myClaims); ok && token2.Valid {
		fmt.Println("Parsed Token:", claims)
	} else {
		fmt.Println("Invalid token")
	}
}

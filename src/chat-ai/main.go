package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"golang.org/x/crypto/bcrypt"
)

// 用户结构体
type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// 登录请求结构体
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// 登录响应结构体
type LoginResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Token   string `json:"token,omitempty"`
}

// 聊天消息结构体
type ChatMessage struct {
	Username string `json:"username"`
	Message  string `json:"message"`
	Time     string `json:"time"`
}

// 用户存储（实际应用中应使用数据库）
var users = make(map[string]User)

// JWT密钥
var jwtKey = []byte("my_secret_key")

// WebSocket升级器
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// WebSocket连接池
var connections = make(map[*websocket.Conn]bool)

func main() {
	router := mux.NewRouter()

	// 静态文件服务
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))

	// API路由
	router.HandleFunc("/login", loginHandler).Methods("POST")
	router.HandleFunc("/register", registerHandler).Methods("POST")
	router.HandleFunc("/chat", chatHandler)

	// 启动服务器
	fmt.Println("服务器启动在 :8080 端口")
	log.Fatal(http.ListenAndServe(":8080", router))
}

// 登录处理
func loginHandler(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "无效的请求", http.StatusBadRequest)
		return
	}

	// 验证用户是否存在
	user, exists := users[req.Username]
	if !exists {
		sendLoginResponse(w, false, "用户不存在", "")
		return
	}

	// 验证密码
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		sendLoginResponse(w, false, "密码错误", "")
		return
	}

	// 生成JWT令牌
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": req.Username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		sendLoginResponse(w, false, "令牌生成失败", "")
		return
	}

	sendLoginResponse(w, true, "登录成功", tokenString)
}

// 注册处理
func registerHandler(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "无效的请求", http.StatusBadRequest)
		return
	}

	// 检查用户是否已存在
	if _, exists := users[req.Username]; exists {
		sendLoginResponse(w, false, "用户已存在", "")
		return
	}

	// 加密密码
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		sendLoginResponse(w, false, "密码加密失败", "")
		return
	}

	// 保存用户
	users[req.Username] = User{
		Username: req.Username,
		Password: string(hashedPassword),
	}

	sendLoginResponse(w, true, "注册成功", "")
}

// 聊天WebSocket处理
func chatHandler(w http.ResponseWriter, r *http.Request) {
	// 升级HTTP连接到WebSocket
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("WebSocket升级失败: %v", err)
		return
	}
	defer conn.Close()

	// 将连接添加到连接池
	connections[conn] = true

	// 处理消息
	for {
		var msg ChatMessage
		err := conn.ReadJSON(&msg)
		if err != nil {
			log.Printf("读取消息错误: %v", err)
			delete(connections, conn)
			break
		}

		// 设置消息时间
		msg.Time = time.Now().Format("15:04:05")

		// 广播消息给所有连接的用户
		broadcastMessage(msg)
	}
}

// 广播消息给所有连接的客户端
func broadcastMessage(msg ChatMessage) {
	for conn := range connections {
		err := conn.WriteJSON(msg)
		if err != nil {
			log.Printf("广播消息错误: %v", err)
			conn.Close()
			delete(connections, conn)
		}
	}
}

// 发送登录响应
func sendLoginResponse(w http.ResponseWriter, success bool, message, token string) {
	response := LoginResponse{
		Success: success,
		Message: message,
		Token:   token,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

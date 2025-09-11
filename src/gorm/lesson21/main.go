package main

import (
	"fmt"

	"database/sql"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 定义模型
type UserInfo struct {
	ID     int64
	Name   *string        `gorm:"default:'unknown'"` //传入空字符串时，默认值不会生效. 有两种方法, 一种是指针类型, 另一种是 sql.NullString
	Gender sql.NullString `gorm:"default:'男'"`       // 另一种是 sql.NullString
	Age    int
}

func main() {
	// 1. 连接数据库
	dsn := "root:root@tcp(120.26.105.128:3306)/db1?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	// 2. 自动迁移模式
	err = db.AutoMigrate(&UserInfo{}) //模型的承载作用, 如果有新字段, 会自动添加到数据库表中
	if err != nil {
		panic("failed to migrate database: " + err.Error())
	}
	// 3. 创建
	name := ""
	user := UserInfo{Name: &name, Age: 25}
	fmt.Println("创建前用户ID:", user.ID)
	db.Create(&user)
	// GORM v2 已移除 NewRecord 方法，可以通过判断主键是否为零值来判断是否为新记录

	db.Debug().First(&user)

	db.FirstOrInit(&user, UserInfo{Age: 30}) // 查找不到就初始化
	fmt.Println("用户ID:", user.ID)
	fmt.Println(user.ID == 0) // true 表示未保存，false 表示已保存

}

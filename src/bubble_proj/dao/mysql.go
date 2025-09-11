package dao

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

func InitMYSQL() error {
	//数据库连接
	dsn := "root:root@tcp(120.26.105.128:3306)/db1?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	//数据库连接池
	sqlDB, err := DB.DB()
	if err != nil {
		return err
	}
	err = sqlDB.Ping()
	if err != nil {
		return err
	}

	//模型绑定
	err = DB.AutoMigrate(&Todo{})
	if err != nil {
		return err
	}

	return nil
}

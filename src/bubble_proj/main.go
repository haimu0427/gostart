package main

import (
	_ "database/sql"
	"proj/dao"
	"proj/router"
)

func main() {
	//数据库
	err := dao.InitMYSQL()
	if err != nil {
		panic(err)
	}
	router := router.SetupRouter()
	router.Run(":8080")
}

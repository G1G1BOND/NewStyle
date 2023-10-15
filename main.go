package main

import (
	"go-v1/database"
	"go-v1/routes"
)

func main() {
	//初始化数据库
	database.InitDB()
	//注册路由
	routes.InitRouter()
}

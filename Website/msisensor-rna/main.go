package main

import (
	"msisensor-rna/model"
	"msisensor-rna/routes"
)

func main() {
	// 初始化数据库
	model.InitDb()

	// 初始化路由
	routes.InitRouter()
}
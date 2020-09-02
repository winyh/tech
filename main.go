package main

import (
	"SJService/router"
)

func main(){
	r := router.InitRouter()
	port := "9090"
	r.Run(":" +  port) // 监听并在 0.0.0.0:8080 上启动服务
}
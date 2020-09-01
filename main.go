package main

import (
	"SJService/router"
)

func main(){
	r := router.InitRouter()
	r.Run() // 监听并在 0.0.0.0:8080 上启动服务
}
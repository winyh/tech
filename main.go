package main

import (
	. "SJService/config"
	"SJService/router"
	"fmt"
)

func main(){
	r := router.InitRouter()
	fmt.Print(PORT)
	r.Run(":" +  PORT) // 监听并在 0.0.0.0:9090 上启动服务
}
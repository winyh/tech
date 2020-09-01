package router

import (
	"SJService/app/Controllers/Home"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	// web路由
	r.GET("/test", Home.Test)

	return r
}

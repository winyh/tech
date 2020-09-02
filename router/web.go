package router

import (
	"SJService/app/Controllers/Admin"
	"SJService/app/Controllers/Home"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	// web路由
	r.GET("/test", Home.Test)
	r.GET("/ping", Admin.Ping)

	// 简单的路由组: v1
	v1 := r.Group("/api")
	{
		v1.POST("/user/create", Admin.UserCreate)
		v1.POST("/user/delete", Admin.UserDestroy)
		v1.POST("/user/update", Admin.UserUpdate)
		v1.GET("/users/:id", Admin.UserFindOne)
		v1.GET("/users", Admin.UserFindAll)
	}

	return r
}

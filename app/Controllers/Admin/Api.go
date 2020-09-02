package Admin

import (
	"SJService/app/Models"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

type Admin struct {
	gorm.Model
	Name string `json:"name"  binding:"required"`
	Password string `json:"password"  binding:"required"`
	Mobile string `json:"mobile" binding:"required"`
}

func Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func UserCreate(c *gin.Context) {

	var json  Models.Admin
	err := c.BindJSON(&json)

	if err != nil {
		fmt.Printf("mysql connect error %v", err)
	}

	id, err := json.Store()

	if err != nil {
		fmt.Printf("database error %v", err)
		fmt.Printf("database error %v", id)
		return
	}

	c.JSON(200, gin.H{
		"status": true,
		"id":id,
		"message":"创建成功",
	})
}

func UserDestroy(c *gin.Context)  {
	var json  Models.Admin
	cid , _ := strconv.Atoi(c.Param("id"))
	err := json.Destroy(cid)

	if err != nil {
		fmt.Printf("database error %v", err)
		c.JSON(200, gin.H{
			"status": false,
			"message":"删除失败",
		})
		return
	}

	c.JSON(200, gin.H{
		"status": true,
		"message":"删除成功",
	})
}

func UserUpdate(c *gin.Context) {
	var json  Models.Admin
	cid , _ := strconv.Atoi(c.Param("id"))
	err := c.BindJSON(&json)

	if err != nil {
		fmt.Printf("mysql connect error %v", err)
	}

	res := json.Update(cid)

	if res != nil {
		fmt.Printf("database error %v", res)
		return
	}

	c.JSON(200, gin.H{
		"status": true,
		"data":cid,
		"message":"更新成功",
	})
}

func UserFindOne(c *gin.Context)  {
	var json  Models.Admin
	id , _ := strconv.Atoi(c.Param("id"))

	result, err := json.FindOne(id)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  false,
			"message": "抱歉未找到相关信息",
		})
		return
	}

	c.JSON(200, gin.H{
		"status": true,
		"data": result,
		"message":"单条查询成功",
	})
}

func UserFindAll(c *gin.Context)  {
	var json  Models.Admin
	err := c.BindJSON(&json)

	if err != nil {
		fmt.Printf("mysql connect error %v", err)
	}

	result, err := json.FindAll()

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  false,
			"message": "抱歉未找到相关信息",
		})
		return
	}

	c.JSON(200, gin.H{
		"status": true,
		"data": result,
		"message":"查询成功",
	})
}
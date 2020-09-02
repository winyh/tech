package Home

import "github.com/gin-gonic/gin"

func Index(c *gin.Context)  {
	c.JSON(200, gin.H{
		"message": "首页测试！",
	})
}

func Test(c *gin.Context)  {
	c.JSON(200, gin.H{
		"message": "首页测试！",
	})
}
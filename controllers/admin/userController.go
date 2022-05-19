package admin

import (
	"github.com/gin-gonic/gin"
)

type UserContoller struct {
}

func (con UserContoller) UserIndex(c *gin.Context) {
	c.HTML(200, "admin/index.html", gin.H{
		"title": "首页+1",
		"score": 80,
		"name":  "zhangsan",
		"hobby": []string{"eat", "sleep", "code"},
		"date":  1652757847,
	})
}

func (con UserContoller) UserAdd(c *gin.Context) {
	c.String(200, "haha")
}

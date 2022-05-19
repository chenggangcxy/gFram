package ruirui

import "github.com/gin-gonic/gin"

//"html/template"
func Index(c *gin.Context) {
	c.HTML(200, "default/index.html", gin.H{
		"title": "首页+1",
		"score": 80,
		"name":  "zhangsan",
		"hobby": []string{"eat", "sleep", "code"},
		"date":  1652757847,
	})
}

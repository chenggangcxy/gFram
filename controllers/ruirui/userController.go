package ruirui

import "github.com/gin-gonic/gin"

type UserController struct {
}

//"html/template"
func (con UserController) Index(c *gin.Context) {
	c.HTML(200, "default/index.html", gin.H{
		"title": "首页+1",
		"score": 80,
		"name":  "zhangsan",
		"hobby": []string{"eat", "sleep", "code"},
		"date":  1652757847,
	})

}

func (con UserController) Useradd(c *gin.Context) {
	c.HTML(200, "default/register.html", gin.H{})
}

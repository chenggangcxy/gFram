package admin

import (
	"path"

	"github.com/gin-gonic/gin"
)

type UserContoller struct {
	BaseController
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

	username := c.PostForm("username")
	//password := c.PostForm("password")

	form, _ := c.MultipartForm()
	files := form.File["face[]"]

	for _, file := range files {
		dst := path.Join("./static/upload", file.Filename)
		c.SaveUploadedFile(file, dst)
	}

	// file, err := c.FormFile("profile")
	// dst := path.Join("./static/upload", file.Filename)
	// if err == nil {

	// 	c.SaveUploadedFile(file, dst)
	// }
	c.JSON(200, gin.H{
		"success":  true,
		"username": username,
	})
}

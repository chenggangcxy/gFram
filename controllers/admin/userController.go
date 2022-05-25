package admin

import (
	"os"
	"path"
	"ruirui/models"
	"strconv"

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

	//form, _ := c.MultipartForm()
	//多文件上传
	file, _ := c.FormFile("profile")

	extname := path.Ext(file.Filename)
	allowExtMap := map[string]bool{
		".jpg":  true,
		".png":  true,
		".gif":  true,
		".jpeg": true,
	}
	if _, ok := allowExtMap[extname]; !ok {
		c.String(200, "上传的文件类型不合法")
		return
	}

	day := models.GetDay()
	dir := "./static/upload/" + day

	arr := os.MkdirAll(dir, 0666)

	if arr != nil {
		c.String(200, "mkdir failed")
	}
	unix := models.GetUnix()

	filename := strconv.FormatInt(unix, 10) + extname

	dst := path.Join(dir, filename)
	c.SaveUploadedFile(file, dst)

	//单文件上传
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

func (con UserContoller) Add(c *gin.Context) {
	userlist := []models.User{}

	models.DB.Find(&userlist)

	//if userlist.username

	c.JSON(200, gin.H{
		"result": userlist,
	})
}

package main

import (
	"ruirui/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	//自制模板函数
	//r.SetFuncMap()
	//加载templates下所有模板
	r.LoadHTMLGlob("templates/**/*")
	//加载静态文件
	r.Static("/static", "./static")

	//
	routers.AdminRouters(r)
	routers.ApiRouters(r)
	routers.DefaultRouters(r)

	r.Run(":80")
}

package main

import (
	"fmt"
	"html/template"
	"ruirui/models"
	"ruirui/routers"
	"time"

	"github.com/gin-gonic/gin"
)

func initMiddleware(c *gin.Context) {
	start := time.Now().UnixNano()
	fmt.Println("midlleware1")
	//next先执行下一个任务，在执行println
	c.Next()
	fmt.Println("middleware2")
	end := time.Now().UnixNano()
	fmt.Println(end - start)

}

//全局中间件

func main() {
	r := gin.Default()
	//自制模板函数
	r.SetFuncMap(template.FuncMap{
		"UnixToTime": models.UnixToTime,
	})

	//加载templates下所有模板
	r.LoadHTMLGlob("templates/**/*")
	//加载静态文件
	r.Static("/static", "./static")
	//全局中间件
	r.Use(initMiddleware)

	routers.AdminRouters(r)
	routers.ApiRouters(r)
	routers.DefaultRouters(r)

	r.Run(":80")
}

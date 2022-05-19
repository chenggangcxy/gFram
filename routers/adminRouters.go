package routers

import (
	"fmt"
	"ruirui/controllers/admin"

	"github.com/gin-gonic/gin"
)

//中间件

func initMiddleware(c *gin.Context) {
	fmt.Println("midlleware1")
	//先执行下一个任务，在执行println
	c.Next()

	fmt.Println("middleware2")

}

func AdminRouters(r *gin.Engine) {
	adminRouters := r.Group("/admin")
	{
		adminRouters.GET("/", initMiddleware, admin.UserContoller{}.UserIndex)
		adminRouters.GET("/useradd", initMiddleware, admin.UserContoller{}.UserAdd)

	}
}

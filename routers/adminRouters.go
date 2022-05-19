package routers

import (
	"fmt"
	"ruirui/controllers/admin"

	"github.com/gin-gonic/gin"
)

//中间件
func initMiddleware(c *gin.Context) {
	fmt.Println("aaa")
}

func AdminRouters(r *gin.Engine) {
	adminRouters := r.Group("/admin")
	{
		adminRouters.GET("/", initMiddleware, admin.UserContoller{}.UserIndex)
		adminRouters.GET("/useradd", admin.UserContoller{}.UserAdd)

	}
}

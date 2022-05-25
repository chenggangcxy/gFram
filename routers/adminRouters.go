package routers

import (
	"ruirui/controllers/admin"
	"ruirui/middlewares"

	"github.com/gin-gonic/gin"
)

func AdminRouters(r *gin.Engine) {
	adminRouters := r.Group("/admin", middlewares.InitMiddleware)
	{
		adminRouters.GET("/", admin.UserContoller{}.UserIndex)
		adminRouters.POST("/user/useradd", admin.UserContoller{}.UserAdd)
		adminRouters.GET("/add", admin.UserContoller{}.Add)

	}
}

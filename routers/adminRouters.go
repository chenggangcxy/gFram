package routers

import (
	"ruirui/controllers/admin"

	"github.com/gin-gonic/gin"
)

func AdminRouters(r *gin.Engine) {
	adminRouters := r.Group("/admin")
	{
		adminRouters.GET("/", admin.UserContoller{}.UserIndex)
		adminRouters.GET("/useradd", admin.UserContoller{}.UserAdd)

	}
}

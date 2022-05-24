package routers

import (
	"ruirui/controllers/ruirui"

	"github.com/gin-gonic/gin"
)

func DefaultRouters(r *gin.Engine) {
	defaultRouters := r.Group("/")
	{
		defaultRouters.GET("/", ruirui.UserController{}.Index)
		defaultRouters.GET("/news", ruirui.NewsContoller{}.Index)
		defaultRouters.GET("/register", ruirui.UserController{}.Useradd)

	}
}

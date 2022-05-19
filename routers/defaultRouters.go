package routers

import (
	"ruirui/controllers/ruirui"

	"github.com/gin-gonic/gin"
)

func DefaultRouters(r *gin.Engine) {
	defaultRouters := r.Group("/")
	{
		defaultRouters.GET("/", ruirui.Index)

	}
}

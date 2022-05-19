package routers

import (
	"ruirui/controllers/api"

	"github.com/gin-gonic/gin"
)

func ApiRouters(r *gin.Engine) {
	apiRouters := r.Group("/api")
	{
		apiRouters.GET("/", api.User)

	}
}

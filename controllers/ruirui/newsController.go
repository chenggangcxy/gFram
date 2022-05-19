package ruirui

import "github.com/gin-gonic/gin"

type NewsContoller struct {
}

func (con NewsContoller) Index(c *gin.Context) {
	c.String(200, "news")
}

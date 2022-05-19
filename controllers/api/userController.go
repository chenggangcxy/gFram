package api

import "github.com/gin-gonic/gin"

type ApiController struct{}

func (con ApiController) User(c *gin.Context) {
	c.HTML(200, "admin/index.html", gin.H{})
}

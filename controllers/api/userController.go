package api

import "github.com/gin-gonic/gin"

func User(c *gin.Context) {
	c.HTML(200, "admin/index.html", gin.H{})
}

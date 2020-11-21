package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// About shows about me page
func About(c *gin.Context) {
	ctx := make(map[string]interface{})
	ctx["Title"] = "关于我"
	c.HTML(http.StatusOK, "blogs/about", ctx)
}

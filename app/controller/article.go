package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/finalone/easy-blog/app/service"
	"github.com/gin-gonic/gin"
)

// ArticleGet handle route /article/:id
func ArticleGet(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		fmt.Println("parse id ", c.Param("id"), "err:", err.Error())
		c.String(http.StatusInternalServerError, err.Error())
	}

	article, err := service.ArticleGetHandler(id)
	if err != nil {
		fmt.Println("get article ", id, "err:", err.Error())
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	c.HTML(http.StatusOK, "article.html", article)
}

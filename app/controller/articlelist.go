package controller

import (
	"fmt"
	"net/http"

	"github.com/finalone/easy-blog/app/service"
	"github.com/gin-gonic/gin"
)

// ArticleListGet handle route to /article_list
func ArticleListGet(c *gin.Context) {
	articles, err := service.ShowArticleList()
	if err != nil {
		fmt.Println("get articles from db err", err)
		c.String(http.StatusInternalServerError, err.Error())
	}

	c.HTML(http.StatusOK, "article_list.html", articles)
}

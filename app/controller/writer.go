package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/finalone/easy-blog/app/model"
	"github.com/finalone/easy-blog/app/service"
	"github.com/gin-gonic/gin"
)

// WriterGet return a html for writer request
func WriterGet(c *gin.Context) {
	idStr := c.Param("id")
	var article *model.Article

	if idStr == "" {
		article = new(model.Article)
		article.Title = "新文章"
	} else {
		id, err := strconv.ParseUint(idStr, 10, 64)
		if err != nil {
			fmt.Println("parse id ", idStr, "err:", err.Error())
			c.String(http.StatusInternalServerError, err.Error())
			return
		}

		article, err = service.ArticleGetHandler(id)
		if err != nil {
			fmt.Println("get article ", id, "err:", err.Error())
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
	}

	c.HTML(http.StatusOK, "writer.html", article)
}

// WriterPost accept a writer post and write to db
func WriterPost(c *gin.Context) {
	var article model.Article
	idStr := c.Param("id")

	article.Title = c.DefaultPostForm("title", "无题")
	article.ClassID, _ = strconv.Atoi(c.DefaultPostForm("classId", "0"))
	article.Abstract.String = c.DefaultPostForm("abstract", "无")
	article.Main.String = c.DefaultPostForm("main", "无")

	var err error
	if idStr != "" {
		id, err := strconv.ParseUint(idStr, 10, 64)
		if err != nil {
			c.String(http.StatusOK, "id错误")
			return
		}
		err = service.UpdateOneArticle(id, &article)
	} else {
		err = service.AddOneArticle(&article)
	}
	if err != nil {
		c.String(http.StatusOK, "编辑失败")
	} else {
		c.String(http.StatusOK, "编辑成功")
	}
}

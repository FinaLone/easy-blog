package main

import (
	"fmt"
	"net/http"

	"github.com/finalone/easy-blog/app/config"
	"github.com/finalone/easy-blog/app/controller"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.Cfg
	fmt.Println(cfg)
	engine := gin.Default()
	engine.LoadHTMLGlob("./template/*")
	engine.Static("/static", "./static")
	engine.Any("/", WebRoot)
	engine.GET("/msgboard", controller.MsgBoardGet)
	engine.POST("/msgboard", controller.MsgBoardPost)
	engine.GET("/article_list", controller.ArticleListGet)
	engine.GET("/article/:id", controller.ArticleGet)
	engine.Run(":12421")
}

func WebRoot(context *gin.Context) {
	context.String(http.StatusOK, "Hello World")
}

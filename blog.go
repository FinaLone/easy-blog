package main

import (
	"fmt"
	"net/http"

	"github.com/finalone/easy-blog/app/config"
	"github.com/finalone/easy-blog/app/controller"
	"github.com/gin-gonic/gin"
	eztemp "github.com/michelloworld/ez-gin-template"
)

func main() {
	cfg := config.Cfg
	fmt.Println(cfg)
	engine := gin.Default()
	render := eztemp.New()
	render.TemplatesDir = "template/"
	render.Debug = true

	engine.HTMLRender = render.Init()

	fmt.Println(render.Templates)
	//engine.LoadHTMLGlob("./template/*")
	engine.Static("/static", "./static")
	engine.Any("/", WebRoot)
	engine.GET("/msgboard", controller.MsgBoardGet)
	engine.POST("/msgboard", controller.MsgBoardPost)
	engine.GET("/article_list", controller.ArticleListGet)
	engine.GET("/article/:id", controller.ArticleGet)
	engine.GET("/writer", controller.WriterGet)
	engine.GET("/writer/:id", controller.WriterGet)
	engine.POST("/writer", controller.WriterPost)
	engine.POST("/writer/:id", controller.WriterPost)
	engine.GET("about", controller.About)
	engine.Run(":12424")
}

func WebRoot(context *gin.Context) {
	context.String(http.StatusOK, "Hello World")
}

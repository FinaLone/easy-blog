package controller

import (
	"fmt"
	"net/http"

	"github.com/finalone/easy-blog/app/model"
	"github.com/finalone/easy-blog/app/service"
	"github.com/gin-gonic/gin"
)

// MsgBoardGet handle GET request of /msgboard
func MsgBoardGet(c *gin.Context) {
	msgs, err := service.MsgBoardGetHandler()
	if err != nil {
		fmt.Println("get msgs from db err", err)
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.HTML(http.StatusOK, "msgboard.html", msgs)
}

// MsgBoardPost handle POST request of /msgboard
func MsgBoardPost(c *gin.Context) {
	var msg model.Message
	msg.AuthorName = c.DefaultPostForm("name", "nobody")
	msg.AuthorEmail = c.DefaultPostForm("mail", "无邮箱信息")
	msg.Msg = c.DefaultPostForm("msg", "")
	service.MsgBoardPostHandler(&msg)
	msgs, err := service.MsgBoardGetHandler()
	if err != nil {
		fmt.Println("get msgs from db err", err)
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.HTML(http.StatusOK, "msgboard.html", msgs)
}

package service

import (
	"github.com/finalone/easy-blog/app/dao"
	"github.com/finalone/easy-blog/app/model"
)

// MsgBoardGetHandler handle GET request to msgboard
func MsgBoardGetHandler() (*[]model.Message, error) {
	//Get msgs from sql
	msgs, err := dao.SelectAllMessage(SqlxDB)
	//fmt.Printf("%t, %v", msgs, msgs)
	if err != nil {
		return nil, err
	}

	return msgs, nil
}

// MsgBoardPostHandler handle POST request to msgboard
func MsgBoardPostHandler(msg *model.Message) {
	//TODO: modify to 异步
	//Write new msg to sql
	dao.InsertOneMessage(SqlxDB, msg)
}

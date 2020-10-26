package dao

import (
	"fmt"

	"github.com/finalone/easy-blog/app/model"
	"github.com/jmoiron/sqlx"
)

// SelectAllMessage get all message from table message_board
func SelectAllMessage(sqlxDB *sqlx.DB) (*[]model.Message, error) {
	var msgs []model.Message
	sql := `SELECT * FROM message_board`
	err := sqlxDB.Select(&msgs, sql)
	if err != nil {
		return nil, err
	}

	return &msgs, nil
}

// InsertOneMessage insert one message into table message
func InsertOneMessage(sqlxDB *sqlx.DB, msg *model.Message) error {
	sql := `INSERT INTO message_board (authorName, authorEmail, message, messageTime) values (?, ?, ?, NOW())`
	_, err := sqlxDB.Exec(sql, msg.AuthorName, msg.AuthorEmail, msg.Msg)
	if err != nil {
		fmt.Println("insert msg err", err)
		return err
	}

	return nil
}

// SelectAllArticleInfo select article main info, to render article_list
func SelectAllArticleInfo(sqlxDB *sqlx.DB) (*[]model.Article, error) {
	var articles []model.Article
	sql := `SELECT id, title, abstract, ctime FROM articles`
	err := sqlxDB.Select(&articles, sql)
	if err != nil {
		return nil, err
	}

	return &articles, nil
}

// SelectArticleDetail select one article detail from table article with id, to render article/:id
func SelectArticleDetail(sqlxDB *sqlx.DB, id uint64) (*model.Article, error) {
	var article model.Article
	sql := `SELECT * FROM articles WHERE id = ?`
	err := sqlxDB.Get(&article, sql, id)
	if err != nil {
		return nil, err
	}
	fmt.Println(article)

	return &article, nil
}

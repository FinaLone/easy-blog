package service

import (
	"github.com/finalone/easy-blog/app/dao"
	"github.com/finalone/easy-blog/app/model"
)

// ShowArticleList process articlelist GET request
func ShowArticleList() (*[]model.Article, error) {
	articles, err := dao.SelectAllArticleInfo(SqlxDB)
	if err != nil {
		return nil, err
	}

	return articles, err
}

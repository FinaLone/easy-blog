package service

import (
	"github.com/finalone/easy-blog/app/dao"
	"github.com/finalone/easy-blog/app/model"
)

// ArticleGetHandler Proccess a request to /article/:id
func ArticleGetHandler(id uint64) (*model.Article, error) {
	article, err := dao.SelectArticleDetail(SqlxDB, id)
	if err != nil {
		return nil, err
	}

	return article, nil
}

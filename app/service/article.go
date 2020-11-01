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

// UpdateOneArticle update you know what
func UpdateOneArticle(id uint64, article *model.Article) error {
	err := dao.UpdateOneArticleByID(SqlxDB, id, article)
	return err
}

// AddOneArticle add you know what
func AddOneArticle(article *model.Article) error {
	err := dao.InsertOneArticle(SqlxDB, article)
	return err
}

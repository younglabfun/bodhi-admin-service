package model

import (
	"gorm.io/gorm"
)

var _ ArticleModel = (*customArticleModel)(nil)

type (
	// ArticleModel is an interface to be customized, add more methods here,
	// and implement the added methods in customArticleModel.
	ArticleModel interface {
		articleModel
	}

	customArticleModel struct {
		*defaultArticleModel
	}
)

// NewArticleModel returns a model for the database table.
func NewArticleModel(conn *gorm.DB) ArticleModel {
	return &customArticleModel{
		defaultArticleModel: newArticleModel(conn),
	}
}

func (m *defaultArticleModel) customCacheKeys(data *Article) []string {
	if data == nil {
		return []string{}
	}
	return []string{}
}

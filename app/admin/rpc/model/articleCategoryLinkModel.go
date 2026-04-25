package model

import (
	"gorm.io/gorm"
)

var _ ArticleCategoryLinkModel = (*customArticleCategoryLinkModel)(nil)

type (
	// ArticleCategoryLinkModel is an interface to be customized, add more methods here,
	// and implement the added methods in customArticleCategoryLinkModel.
	ArticleCategoryLinkModel interface {
		articleCategoryLinkModel
	}

	customArticleCategoryLinkModel struct {
		*defaultArticleCategoryLinkModel
	}
)

// NewArticleCategoryLinkModel returns a model for the database table.
func NewArticleCategoryLinkModel(conn *gorm.DB) ArticleCategoryLinkModel {
	return &customArticleCategoryLinkModel{
		defaultArticleCategoryLinkModel: newArticleCategoryLinkModel(conn),
	}
}

func (m *defaultArticleCategoryLinkModel) customCacheKeys(data *ArticleCategoryLink) []string {
	if data == nil {
		return []string{}
	}
	return []string{}
}

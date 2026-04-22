package model

import (
	"gorm.io/gorm"
)

var _ MediaModel = (*customMediaModel)(nil)

type (
	// MediaModel is an interface to be customized, add more methods here,
	// and implement the added methods in customMediaModel.
	MediaModel interface {
		mediaModel
	}

	customMediaModel struct {
		*defaultMediaModel
	}
)

// NewMediaModel returns a model for the database table.
func NewMediaModel(conn *gorm.DB) MediaModel {
	return &customMediaModel{
		defaultMediaModel: newMediaModel(conn),
	}
}

func (m *defaultMediaModel) customCacheKeys(data *Media) []string {
	if data == nil {
		return []string{}
	}
	return []string{}
}

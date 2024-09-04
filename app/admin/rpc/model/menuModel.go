package model

import (
	"gorm.io/gorm"
)

var _ MenuModel = (*customMenuModel)(nil)

type (
	// MenuModel is an interface to be customized, add more methods here,
	// and implement the added methods in customMenuModel.
	MenuModel interface {
		menuModel
	}

	customMenuModel struct {
		*defaultMenuModel
	}
)

// NewMenuModel returns a model for the database table.
func NewMenuModel(conn *gorm.DB) MenuModel {
	return &customMenuModel{
		defaultMenuModel: newMenuModel(conn),
	}
}

func (m *defaultMenuModel) customCacheKeys(data *Menu) []string {
	if data == nil {
		return []string{}
	}
	return []string{}
}

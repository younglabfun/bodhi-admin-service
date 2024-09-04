package model

import (
	"gorm.io/gorm"
)

var _ UserTokenModel = (*customUserTokenModel)(nil)

type (
	// UserTokenModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserTokenModel.
	UserTokenModel interface {
		userTokenModel
	}

	customUserTokenModel struct {
		*defaultUserTokenModel
	}
)

// NewUserTokenModel returns a model for the database table.
func NewUserTokenModel(conn *gorm.DB) UserTokenModel {
	return &customUserTokenModel{
		defaultUserTokenModel: newUserTokenModel(conn),
	}
}

func (m *defaultUserTokenModel) customCacheKeys(data *UserToken) []string {
	if data == nil {
		return []string{}
	}
	return []string{}
}

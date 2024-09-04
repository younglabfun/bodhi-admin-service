package model

import (
	"gorm.io/gorm"
)

var _ UserRoleModel = (*customUserRoleModel)(nil)

type (
	// UserRoleModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserRoleModel.
	UserRoleModel interface {
		userRoleModel
	}

	customUserRoleModel struct {
		*defaultUserRoleModel
	}
)

// NewUserRoleModel returns a model for the database table.
func NewUserRoleModel(conn *gorm.DB) UserRoleModel {
	return &customUserRoleModel{
		defaultUserRoleModel: newUserRoleModel(conn),
	}
}

func (m *defaultUserRoleModel) customCacheKeys(data *UserRole) []string {
	if data == nil {
		return []string{}
	}
	return []string{}
}

package model

import (
	"gorm.io/gorm"
)

var _ NodeGroupModel = (*customNodeGroupModel)(nil)

type (
	// NodeGroupModel is an interface to be customized, add more methods here,
	// and implement the added methods in customNodeGroupModel.
	NodeGroupModel interface {
		nodeGroupModel
	}

	customNodeGroupModel struct {
		*defaultNodeGroupModel
	}
)

// NewNodeGroupModel returns a model for the database table.
func NewNodeGroupModel(conn *gorm.DB) NodeGroupModel {
	return &customNodeGroupModel{
		defaultNodeGroupModel: newNodeGroupModel(conn),
	}
}

func (m *defaultNodeGroupModel) customCacheKeys(data *NodeGroup) []string {
	if data == nil {
		return []string{}
	}
	return []string{}
}

package model

import (
	"gorm.io/gorm"
)

var _ NodeModel = (*customNodeModel)(nil)

type (
	// NodeModel is an interface to be customized, add more methods here,
	// and implement the added methods in customNodeModel.
	NodeModel interface {
		nodeModel
	}

	customNodeModel struct {
		*defaultNodeModel
	}
)

// NewNodeModel returns a model for the database table.
func NewNodeModel(conn *gorm.DB) NodeModel {
	return &customNodeModel{
		defaultNodeModel: newNodeModel(conn),
	}
}

func (m *defaultNodeModel) customCacheKeys(data *Node) []string {
	if data == nil {
		return []string{}
	}
	return []string{}
}

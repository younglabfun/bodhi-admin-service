// Code generated by goctl. DO NOT EDIT!

package model

import (
	"bodhiadmin/common/utils"
	"context"
	"github.com/SpectatorNan/gorm-zero/gormc"
	"gorm.io/gorm"
	"gorm.io/plugin/soft_delete"
)

var ()

type (
	nodeModel interface {
		Insert(ctx context.Context, data *Node) error
		FindOne(ctx context.Context, id int64) (*Node, error)
		FindOneByFuncCode(ctx context.Context, funcCode string) (*Node, error)
		Update(ctx context.Context, data *Node) error
		Delete(ctx context.Context, id int64) error

		FindAll(ctx context.Context) ([]*Node, error)
		FindListByGid(ctx context.Context, groupId int64) ([]*Node, error)
		FindListByIds(ctx context.Context, ids []int64) ([]*Node, error)
		FindListByPage(ctx context.Context, req PageReq) ([]*Node, int64, error)
		BatchUpdateNodeGroup(ctx context.Context, ids []int64, groupId int64) error
		UpdateStatus(ctx context.Context, id int64) error
		BatchDelete(ctx context.Context, ids []int64) error
	}

	defaultNodeModel struct {
		conn  *gorm.DB
		table string
	}

	Node struct {
		Id          int64                 `gorm:"column:id;primaryKey;unique"` // 功能ID
		GroupId     int64                 `gorm:"column:group_id"`             // 应用功能分组ID
		FuncCode    string                `gorm:"column:func_code"`            // 功能标识
		Name        string                `gorm:"column:name"`                 // 功能名称
		Description string                `gorm:"column:description"`          // 功能描述
		IsEnabled   int64                 `gorm:"default:1"`                   // 是否启用
		IsDeleted   soft_delete.DeletedAt `gorm:"softDelete:flag,default:0"`   // 是否删除
		CreatedAt   int64                 `gorm:"<-:create;autoCreateTime"`    // 添加时间
	}
)

func (Node) TableName() string {
	return "`node`"
}

func newNodeModel(conn *gorm.DB) *defaultNodeModel {
	return &defaultNodeModel{
		conn:  conn,
		table: "`node`",
	}
}

func (m *defaultNodeModel) Insert(ctx context.Context, data *Node) error {
	err := m.conn.WithContext(ctx).Save(&data).Error
	return err
}

func (m *defaultNodeModel) FindOne(ctx context.Context, id int64) (*Node, error) {
	var resp Node
	err := m.conn.WithContext(ctx).Model(&Node{}).Where("`id` = ?", id).Take(&resp).Error
	switch err {
	case nil:
		return &resp, nil
	case gormc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultNodeModel) FindOneByFuncCode(ctx context.Context, funcCode string) (*Node, error) {
	var resp Node
	err := m.conn.WithContext(ctx).Model(&Node{}).Where("`func_code` = ?", funcCode).Take(&resp).Error
	switch err {
	case nil:
		return &resp, nil
	case gormc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultNodeModel) Update(ctx context.Context, data *Node) error {
	err := m.conn.WithContext(ctx).Omit("is_enabled").Save(data).Error
	return err
}

func (m *defaultNodeModel) Delete(ctx context.Context, id int64) error {
	err := m.conn.WithContext(ctx).Delete(&Node{}, id).Error

	return err
}

func (m *defaultNodeModel) tableName() string {
	return m.table
}

func (m *defaultNodeModel) FindAll(ctx context.Context) ([]*Node, error) {
	var list []*Node
	db := m.conn.WithContext(ctx).Model(Node{})
	db.Order("group_id ASC")
	err := db.Find(&list).Error

	if err != nil {
		return nil, nil
	}
	return list, nil
}

func (m *defaultNodeModel) FindListByGid(ctx context.Context, groupId int64) ([]*Node, error) {
	var list []*Node
	db := m.conn.WithContext(ctx).Model(&Node{})
	db.Where("group_id = ?", groupId)
	db.Order("created_at DESC")
	err := db.Find(&list).Error

	if err != nil {
		return nil, err
	}
	return list, nil
}

func (m *defaultNodeModel) FindListByIds(ctx context.Context, ids []int64) ([]*Node, error) {
	var list []*Node
	db := m.conn.WithContext(ctx).Model(&Node{})
	db.Where("id in ?", ids)
	err := db.Find(&list).Error

	if err != nil {
		return nil, err
	}
	return list, nil
}

func (m *defaultNodeModel) FindListByPage(ctx context.Context, req PageReq) ([]*Node, int64, error) {
	var list []*Node
	var total int64

	sortBy := utils.GetSortByStr(req.Sort, req.Order, User{})
	db := m.conn.WithContext(ctx).Model(Node{})
	if len(req.Field) != 0 && req.Field == "group_id" && len(req.Value) != 0 {
		db.Where("`group_id` = ?", req.Value)
	}
	db.Count(&total)
	db.Scopes(utils.Paginate(req.Page, req.Size))
	db.Order(sortBy)
	err := db.Find(&list).Error

	if err != nil {
		return nil, 0, err
	}
	return list, total, nil
}

func (m *defaultNodeModel) BatchUpdateNodeGroup(ctx context.Context, ids []int64, groupId int64) error {
	db := m.conn.WithContext(ctx).Model(Node{})
	db.Where("id IN ?", ids)
	err := db.Update("GroupId", groupId).Error

	return err
}

func (m *defaultNodeModel) UpdateStatus(ctx context.Context, id int64) error {
	old, err := m.FindOne(ctx, id)
	if err != nil {
		return err
	}

	data := 0
	if old.IsEnabled == 0 {
		data = 1
	}

	db := m.conn.WithContext(ctx).Model(&Node{})
	err = db.Where("id = ?", id).Update("IsEnabled", data).Error
	return err
}

func (m *defaultNodeModel) BatchDelete(ctx context.Context, ids []int64) error {
	_, err := m.FindListByIds(ctx, ids)
	if err != nil {
		return err
	}
	err = m.conn.WithContext(ctx).Delete(&Node{}, ids).Error
	return err
}

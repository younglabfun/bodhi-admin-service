package svc

import (
	"bodhiadmin/app/admin/rpc/internal/config"
	"bodhiadmin/app/admin/rpc/model"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

type ServiceContext struct {
	Config config.Config

	MenuModel      model.MenuModel
	NodeModel      model.NodeModel
	NodeGroupModel model.NodeGroupModel
	RoleModel      model.RoleModel
	UserModel      model.UserModel
	UserRoleModel  model.UserRoleModel
	TokenModel     model.UserTokenModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	db, err := gorm.Open(mysql.Open(c.DB.DataSource), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 使用单数表名，启用该选项后，`User` 表将是`user`
		},
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		logx.Error("Init Mysql connect error:", err)
		panic(any(err))
	}

	menuModel := model.NewMenuModel(db)
	nodeModel := model.NewNodeModel(db)
	nodeGroupModel := model.NewNodeGroupModel(db)

	roleModel := model.NewRoleModel(db)
	userModel := model.NewUserModel(db)
	userRoleModel := model.NewUserRoleModel(db)
	tokenModel := model.NewUserTokenModel(db)

	return &ServiceContext{
		Config: c,

		MenuModel:      menuModel,
		NodeModel:      nodeModel,
		NodeGroupModel: nodeGroupModel,
		RoleModel:      roleModel,
		UserModel:      userModel,
		UserRoleModel:  userRoleModel,
		TokenModel:     tokenModel,
	}
}

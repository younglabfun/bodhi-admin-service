package svc

import (
	"bodhiadmin/app/admin/rpc/internal/config"
	"bodhiadmin/app/admin/rpc/model"
	"bodhiadmin/common/utils"
	"github.com/zeromicro/go-zero/core/logx"
)

type ServiceContext struct {
	Config config.Config

	MediaModel     model.MediaModel
	MenuModel      model.MenuModel
	NodeModel      model.NodeModel
	NodeGroupModel model.NodeGroupModel
	RoleModel      model.RoleModel
	UserModel      model.UserModel
	UserRoleModel  model.UserRoleModel
	TokenModel     model.UserTokenModel
	CategoryModel  model.CategoryModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	db, err := utils.InitMySQL(c.MySql.Host, c.MySql.User, c.MySql.Password, c.MySql.Database, c.MySql.Port)
	if err != nil {
		logx.Error("Init MySQL connect error:", err)
		panic(any(err))
	}
	logx.Infof("Init MySQL connected...")

	mediaModel := model.NewMediaModel(db)
	categoryModel := model.NewCategoryModel(db)
	menuModel := model.NewMenuModel(db)
	nodeModel := model.NewNodeModel(db)
	nodeGroupModel := model.NewNodeGroupModel(db)

	roleModel := model.NewRoleModel(db)
	userModel := model.NewUserModel(db)
	userRoleModel := model.NewUserRoleModel(db)
	tokenModel := model.NewUserTokenModel(db)

	return &ServiceContext{
		Config: c,

		MediaModel:     mediaModel,
		CategoryModel:  categoryModel,
		MenuModel:      menuModel,
		NodeModel:      nodeModel,
		NodeGroupModel: nodeGroupModel,
		RoleModel:      roleModel,
		UserModel:      userModel,
		UserRoleModel:  userRoleModel,
		TokenModel:     tokenModel,
	}
}

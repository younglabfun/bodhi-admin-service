package accountlogic

import (
	"bodhiadmin/app/admin/rpc/internal/svc"
	"bodhiadmin/app/admin/rpc/model"
	"bodhiadmin/app/admin/rpc/proto/admin"
	"context"
	"encoding/json"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPermissionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetPermissionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPermissionLogic {
	return &GetPermissionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetPermissionLogic) GetPermission(in *admin.Uuid) (*admin.PermissionResp, error) {
	var permission []string
	var nodes []*model.Node
	var err error

	if in.Uuid == l.svcCtx.Config.AdminConf.Master {
		nodes, err = l.svcCtx.NodeModel.FindAll(l.ctx)
		if err != nil {
			return nil, err
		}

	} else {
		roles, err := l.svcCtx.UserRoleModel.FindUserRoles(l.ctx, in.Uuid)
		if err != nil {
			return nil, err
		}

		//fmt.Println("user roles", len(resp))
		if len(roles) != 0 {
			var authIds []int64
			for _, v := range roles {
				var roleAuthData []int64
				json.Unmarshal([]byte(v.Role.AuthorizeJson), &roleAuthData)

				authIds = append(authIds, roleAuthData...)
			}
			if len(authIds) != 0 {
				nodes, err = l.svcCtx.NodeModel.FindListByIds(l.ctx, authIds)
				if err != nil {
					return nil, err
				}
			}
		}
	}

	if len(nodes) != 0 {
		for _, n := range nodes {
			permission = append(permission, n.FuncCode)
		}
	}
	//fmt.Println("permission", len(permission))
	return &admin.PermissionResp{
		Permission: permission,
	}, nil
}

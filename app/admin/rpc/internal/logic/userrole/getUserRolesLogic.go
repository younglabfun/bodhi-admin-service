package userrolelogic

import (
	"context"
	"github.com/jinzhu/copier"

	"bodhiadmin/app/admin/rpc/internal/svc"
	"bodhiadmin/app/admin/rpc/proto/admin"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserRolesLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserRolesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserRolesLogic {
	return &GetUserRolesLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserRolesLogic) GetUserRoles(in *admin.Uuid) (*admin.UserRoleListResp, error) {
	resp, err := l.svcCtx.UserRoleModel.FindUserRoles(l.ctx, in.Uuid)
	if err != nil {
		return nil, err
	}

	var list []*admin.UserRoleUnit
	if len(resp) != 0 {
		for _, v := range resp {
			var item admin.UserRoleUnit
			_ = copier.Copy(&item, v)

			list = append(list, &item)
		}
	}

	return &admin.UserRoleListResp{
		List: list,
	}, nil
}

package user

import (
	"bodhiadmin/app/admin/rpc/proto/admin"
	"context"
	"github.com/jinzhu/copier"

	"bodhiadmin/app/admin/api/internal/svc"
	"bodhiadmin/app/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserRolesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserRolesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserRolesLogic {
	return &GetUserRolesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserRolesLogic) GetUserRoles(req *types.UuidPath) (*types.UserRolesResp, error) {
	resp, err := l.svcCtx.UserRoleRpc.GetUserRoles(l.ctx, &admin.Uuid{
		Uuid: req.Uuid,
	})
	if err != nil {
		return nil, err
	}
	var list []*types.UserRoleUnit
	if len(resp.List) != 0 {
		for _, v := range resp.List {
			var item types.UserRoleUnit
			_ = copier.Copy(&item, v)

			list = append(list, &item)
		}
	}

	return &types.UserRolesResp{
		List: list,
	}, nil
}

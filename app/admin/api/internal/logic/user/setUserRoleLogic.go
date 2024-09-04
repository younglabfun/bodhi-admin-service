package user

import (
	"bodhiadmin/app/admin/rpc/proto/admin"
	"context"

	"bodhiadmin/app/admin/api/internal/svc"
	"bodhiadmin/app/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SetUserRoleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSetUserRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetUserRoleLogic {
	return &SetUserRoleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SetUserRoleLogic) SetUserRole(req *types.UserRoleReq) (*types.AffectedResp, error) {
	resp, err := l.svcCtx.UserRoleRpc.InsertUserRole(l.ctx, &admin.UserRoleReq{
		UserUuid: req.UserUuid,
		RoleUuid: req.RoleUuid,
	})
	if err != nil {
		return nil, err
	}

	return &types.AffectedResp{
		Affected: resp.Affected,
	}, nil
}

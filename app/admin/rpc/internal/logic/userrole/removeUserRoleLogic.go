package userrolelogic

import (
	"context"

	"bodhiadmin/app/admin/rpc/internal/svc"
	"bodhiadmin/app/admin/rpc/proto/admin"

	"github.com/zeromicro/go-zero/core/logx"
)

type RemoveUserRoleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRemoveUserRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RemoveUserRoleLogic {
	return &RemoveUserRoleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RemoveUserRoleLogic) RemoveUserRole(in *admin.Id) (*admin.AffectedResp, error) {
	err := l.svcCtx.UserRoleModel.Delete(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}

	return &admin.AffectedResp{
		Affected: true,
	}, nil
}

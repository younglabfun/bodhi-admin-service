package rolelogic

import (
	"context"

	"bodhiadmin/app/admin/rpc/internal/svc"
	"bodhiadmin/app/admin/rpc/proto/admin"

	"github.com/zeromicro/go-zero/core/logx"
)

type RemoveRoleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRemoveRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RemoveRoleLogic {
	return &RemoveRoleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RemoveRoleLogic) RemoveRole(in *admin.Uuid) (*admin.AffectedResp, error) {
	err := l.svcCtx.RoleModel.Delete(l.ctx, in.Uuid)
	if err != nil {
		return nil, err
	}

	return &admin.AffectedResp{
		Affected: true,
	}, nil
}

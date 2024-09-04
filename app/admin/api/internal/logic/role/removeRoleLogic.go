package role

import (
	"bodhiadmin/app/admin/rpc/proto/admin"
	"context"

	"bodhiadmin/app/admin/api/internal/svc"
	"bodhiadmin/app/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RemoveRoleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRemoveRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RemoveRoleLogic {
	return &RemoveRoleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RemoveRoleLogic) RemoveRole(req *types.UuidReq) (*types.AffectedResp, error) {
	resp, err := l.svcCtx.RoleRpc.RemoveRole(l.ctx, &admin.Uuid{
		Uuid: req.Uuid,
	})
	if err != nil {
		return nil, err
	}

	return &types.AffectedResp{
		Affected: resp.Affected,
	}, nil
}

package role

import (
	"bodhiadmin/app/admin/rpc/proto/admin"
	"context"

	"bodhiadmin/app/admin/api/internal/svc"
	"bodhiadmin/app/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateStatusLogic {
	return &UpdateStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateStatusLogic) UpdateStatus(req *types.UuidStatusReq) (*types.AffectedResp, error) {
	resp, err := l.svcCtx.RoleRpc.UpdateStatus(l.ctx, &admin.UuidStatusReq{
		Uuid:   req.Uuid,
		Status: req.Status,
	})
	if err != nil {
		return nil, err
	}

	return &types.AffectedResp{
		Affected: resp.Affected,
	}, nil
}

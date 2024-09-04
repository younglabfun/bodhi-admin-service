package node

import (
	"bodhiadmin/app/admin/rpc/proto/admin"
	"context"

	"bodhiadmin/app/admin/api/internal/svc"
	"bodhiadmin/app/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateNodeStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateNodeStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateNodeStatusLogic {
	return &UpdateNodeStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateNodeStatusLogic) UpdateNodeStatus(req *types.IdReq) (*types.AffectedResp, error) {
	resp, err := l.svcCtx.NodeRpc.UpdateStatus(l.ctx, &admin.Id{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}

	return &types.AffectedResp{
		Affected: resp.Affected,
	}, nil
}

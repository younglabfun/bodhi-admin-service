package nodeGroup

import (
	"bodhiadmin/app/admin/rpc/proto/admin"
	"context"

	"bodhiadmin/app/admin/api/internal/svc"
	"bodhiadmin/app/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RemoveGroupLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRemoveGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RemoveGroupLogic {
	return &RemoveGroupLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RemoveGroupLogic) RemoveGroup(req *types.IdReq) (*types.AffectedResp, error) {
	resp, err := l.svcCtx.NodeGroupRpc.RemoveNodeGroup(l.ctx, &admin.Id{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}

	return &types.AffectedResp{
		Affected: resp.Affected,
	}, nil
}

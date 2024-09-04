package node

import (
	"bodhiadmin/app/admin/rpc/proto/admin"
	"context"

	"bodhiadmin/app/admin/api/internal/svc"
	"bodhiadmin/app/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RemoveNodeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRemoveNodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RemoveNodeLogic {
	return &RemoveNodeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RemoveNodeLogic) RemoveNode(req *types.IdReq) (resp *types.AffectedResp, err error) {
	r, e := l.svcCtx.NodeRpc.RemoveNode(l.ctx, &admin.Id{
		Id: req.Id,
	})
	if e != nil {
		return nil, e
	}

	return &types.AffectedResp{
		Affected: r.Affected,
	}, nil
}

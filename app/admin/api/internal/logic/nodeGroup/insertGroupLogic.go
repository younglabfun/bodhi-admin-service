package nodeGroup

import (
	"bodhiadmin/app/admin/rpc/proto/admin"
	"context"

	"bodhiadmin/app/admin/api/internal/svc"
	"bodhiadmin/app/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type InsertGroupLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewInsertGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InsertGroupLogic {
	return &InsertGroupLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *InsertGroupLogic) InsertGroup(req *types.NodeGroupReq) (*types.AffectedResp, error) {
	nodeGroup := admin.NodeGroupReq{
		Name:  req.Name,
		Title: req.Title,
		Sort:  req.Sort,
	}
	resp, err := l.svcCtx.NodeGroupRpc.InsertNodeGroup(l.ctx, &nodeGroup)
	if err != nil {
		return nil, err
	}

	return &types.AffectedResp{
		Affected: resp.Affected,
	}, nil
}

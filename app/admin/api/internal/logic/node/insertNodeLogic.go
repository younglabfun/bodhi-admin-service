package node

import (
	"bodhiadmin/app/admin/rpc/proto/admin"
	"context"

	"bodhiadmin/app/admin/api/internal/svc"
	"bodhiadmin/app/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type InsertNodeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewInsertNodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InsertNodeLogic {
	return &InsertNodeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *InsertNodeLogic) InsertNode(req *types.NodeReq) (*types.AffectedResp, error) {
	node := admin.NodeReq{
		GroupId:     req.GroupId,
		FuncCode:    req.FuncCode,
		Name:        req.Name,
		Description: req.Description,
	}
	resp, err := l.svcCtx.NodeRpc.InsertNode(l.ctx, &node)
	if err != nil {
		return nil, err
	}

	return &types.AffectedResp{
		Affected: resp.Affected,
	}, nil
}

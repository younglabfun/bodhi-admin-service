package node

import (
	"bodhiadmin/app/admin/rpc/proto/admin"
	"context"
	"github.com/jinzhu/copier"

	"bodhiadmin/app/admin/api/internal/svc"
	"bodhiadmin/app/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetNodeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetNodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetNodeLogic {
	return &GetNodeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetNodeLogic) GetNode(req *types.IdPath) (*types.NodeResp, error) {
	resp, err := l.svcCtx.NodeRpc.GetNode(l.ctx, &admin.Id{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}
	var node types.NodeResp
	_ = copier.Copy(&node, resp)

	return &node, nil
}

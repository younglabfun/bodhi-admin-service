package node

import (
	"bodhiadmin/app/admin/rpc/proto/admin"
	"context"

	"bodhiadmin/app/admin/api/internal/svc"
	"bodhiadmin/app/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateNodeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateNodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateNodeLogic {
	return &UpdateNodeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateNodeLogic) UpdateNode(req *types.NodeReq) (*types.AffectedResp, error) {
	node := admin.NodeReq{
		Id:          req.Id,
		GroupId:     req.GroupId,
		FuncCode:    req.FuncCode,
		Name:        req.Name,
		Description: req.Description,
	}

	resp, err := l.svcCtx.NodeRpc.UpdateNode(l.ctx, &node)
	if err != nil {
		return nil, err
	}

	return &types.AffectedResp{
		Affected: resp.Affected,
	}, nil
}

package nodeGroup

import (
	"bodhiadmin/app/admin/rpc/proto/admin"
	"context"

	"bodhiadmin/app/admin/api/internal/svc"
	"bodhiadmin/app/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateGroupLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateGroupLogic {
	return &UpdateGroupLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateGroupLogic) UpdateGroup(req *types.NodeGroupReq) (*types.AffectedResp, error) {
	nodeGroup := admin.NodeGroupReq{
		Id:    req.Id,
		Name:  req.Name,
		Title: req.Title,
		Sort:  req.Sort,
	}
	resp, err := l.svcCtx.NodeGroupRpc.UpdateNodeGroup(l.ctx, &nodeGroup)
	if err != nil {
		return nil, err
	}

	return &types.AffectedResp{
		Affected: resp.Affected,
	}, nil
}

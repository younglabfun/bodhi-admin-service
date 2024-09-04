package nodelogic

import (
	"context"
	"github.com/jinzhu/copier"

	"bodhiadmin/app/admin/rpc/internal/svc"
	"bodhiadmin/app/admin/rpc/proto/admin"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetNodeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetNodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetNodeLogic {
	return &GetNodeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetNodeLogic) GetNode(in *admin.Id) (*admin.NodeUnit, error) {
	resp, err := l.svcCtx.NodeModel.FindOne(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}

	var node admin.NodeUnit
	_ = copier.Copy(&node, resp)

	return &node, nil
}

package nodegrouplogic

import (
	"context"
	"github.com/jinzhu/copier"

	"bodhiadmin/app/admin/rpc/internal/svc"
	"bodhiadmin/app/admin/rpc/proto/admin"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetNodeGroupLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetNodeGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetNodeGroupLogic {
	return &GetNodeGroupLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetNodeGroupLogic) GetNodeGroup(in *admin.Id) (*admin.NodeGroupUnit, error) {
	resp, err := l.svcCtx.NodeGroupModel.FindOne(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}
	var group admin.NodeGroupUnit
	_ = copier.Copy(&group, resp)

	return &group, nil
}

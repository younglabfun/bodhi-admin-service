package nodelogic

import (
	"context"

	"bodhiadmin/app/admin/rpc/internal/svc"
	"bodhiadmin/app/admin/rpc/proto/admin"

	"github.com/zeromicro/go-zero/core/logx"
)

type RemoveNodeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRemoveNodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RemoveNodeLogic {
	return &RemoveNodeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RemoveNodeLogic) RemoveNode(in *admin.Id) (*admin.AffectedResp, error) {
	err := l.svcCtx.NodeModel.Delete(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}

	return &admin.AffectedResp{
		Affected: true,
	}, nil
}

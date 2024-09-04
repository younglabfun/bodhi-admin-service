package nodelogic

import (
	"context"
	"github.com/pkg/errors"

	"bodhiadmin/app/admin/rpc/internal/svc"
	"bodhiadmin/app/admin/rpc/proto/admin"

	"github.com/zeromicro/go-zero/core/logx"
)

type BatchRemoveLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewBatchRemoveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BatchRemoveLogic {
	return &BatchRemoveLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *BatchRemoveLogic) BatchRemove(in *admin.BatchIdsReq) (*admin.AffectedResp, error) {
	if len(in.Ids) == 0 {
		return nil, errors.New("params is invalid")
	}
	err := l.svcCtx.NodeModel.BatchDelete(l.ctx, in.Ids)
	if err != nil {
		return nil, err
	}

	return &admin.AffectedResp{
		Affected: true,
	}, nil
}

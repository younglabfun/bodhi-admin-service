package nodelogic

import (
	"context"
	"github.com/pkg/errors"

	"bodhiadmin/app/admin/rpc/internal/svc"
	"bodhiadmin/app/admin/rpc/proto/admin"

	"github.com/zeromicro/go-zero/core/logx"
)

type BatchMoveLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewBatchMoveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BatchMoveLogic {
	return &BatchMoveLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *BatchMoveLogic) BatchMove(in *admin.MoveReq) (*admin.AffectedResp, error) {
	if len(in.Ids) == 0 {
		return nil, errors.New("params is invalid")
	}
	_, err := l.svcCtx.NodeGroupModel.FindOne(l.ctx, in.GroupId)
	if err != nil {
		return nil, err
	}

	err = l.svcCtx.NodeModel.BatchUpdateNodeGroup(l.ctx, in.Ids, in.GroupId)
	if err != nil {
		return nil, err
	}

	return &admin.AffectedResp{
		Affected: true,
	}, nil
}

package node

import (
	"bodhiadmin/app/admin/rpc/proto/admin"
	"context"
	"encoding/json"
	"github.com/pkg/errors"

	"bodhiadmin/app/admin/api/internal/svc"
	"bodhiadmin/app/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type BatchMoveGroupLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewBatchMoveGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BatchMoveGroupLogic {
	return &BatchMoveGroupLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *BatchMoveGroupLogic) BatchMoveGroup(req *types.MoveReq) (*types.AffectedResp, error) {
	var ids []int64
	err := json.Unmarshal([]byte(req.Ids), &ids)
	if err != nil {
		return nil, err
	}
	if len(ids) == 0 {
		return nil, errors.New("params is invalid")
	}

	resp, err := l.svcCtx.NodeRpc.BatchMove(l.ctx, &admin.MoveReq{
		GroupId: req.GroupId,
		Ids:     ids,
	})
	if err != nil {
		return nil, err
	}

	return &types.AffectedResp{
		Affected: resp.Affected,
	}, nil
}

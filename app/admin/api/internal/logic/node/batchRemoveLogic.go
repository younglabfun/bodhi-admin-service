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

type BatchRemoveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewBatchRemoveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BatchRemoveLogic {
	return &BatchRemoveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *BatchRemoveLogic) BatchRemove(req *types.BatchRemoveReq) (*types.AffectedResp, error) {
	var ids []int64
	err := json.Unmarshal([]byte(req.Ids), &ids)
	if err != nil {
		return nil, err
	}
	if len(ids) == 0 {
		return nil, errors.New("params is invalid")
	}

	resp, err := l.svcCtx.NodeRpc.BatchRemove(l.ctx, &admin.BatchIdsReq{
		Ids: ids,
	})
	if err != nil {
		return nil, err
	}

	return &types.AffectedResp{
		Affected: resp.Affected,
	}, nil
}

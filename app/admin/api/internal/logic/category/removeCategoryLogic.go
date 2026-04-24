package category

import (
	"bodhiadmin/app/admin/rpc/proto/admin"
	"context"

	"bodhiadmin/app/admin/api/internal/svc"
	"bodhiadmin/app/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RemoveCategoryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRemoveCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RemoveCategoryLogic {
	return &RemoveCategoryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RemoveCategoryLogic) RemoveCategory(req *types.IdReq) (*types.AffectedResp, error) {
	resp, err := l.svcCtx.CategoryRpc.RemoveCategory(l.ctx, &admin.Id{
		Id: req.Id,
	})

	return &types.AffectedResp{
		Affected: resp.Affected,
	}, err
}

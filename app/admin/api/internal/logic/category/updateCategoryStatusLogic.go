package category

import (
	"bodhiadmin/app/admin/rpc/proto/admin"
	"context"

	"bodhiadmin/app/admin/api/internal/svc"
	"bodhiadmin/app/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateCategoryStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateCategoryStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateCategoryStatusLogic {
	return &UpdateCategoryStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateCategoryStatusLogic) UpdateCategoryStatus(req *types.IdReq) (*types.AffectedResp, error) {
	resp, err := l.svcCtx.CategoryRpc.UpdateStatus(l.ctx, &admin.Id{
		Id: req.Id,
	})

	return &types.AffectedResp{
		Affected: resp.Affected,
	}, err
}

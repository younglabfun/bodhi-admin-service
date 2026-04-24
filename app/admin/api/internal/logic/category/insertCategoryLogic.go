package category

import (
	"bodhiadmin/app/admin/rpc/proto/admin"
	"context"
	"github.com/jinzhu/copier"

	"bodhiadmin/app/admin/api/internal/svc"
	"bodhiadmin/app/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type InsertCategoryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewInsertCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InsertCategoryLogic {
	return &InsertCategoryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *InsertCategoryLogic) InsertCategory(req *types.CategoryReq) (*types.AffectedResp, error) {
	var data admin.CategoryReq
	_ = copier.Copy(&data, req)
	resp, err := l.svcCtx.CategoryRpc.InsertCategory(l.ctx, &data)

	return &types.AffectedResp{
		Affected: resp.Affected,
	}, err
}

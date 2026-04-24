package category

import (
	"bodhiadmin/app/admin/rpc/proto/admin"
	"context"
	"github.com/jinzhu/copier"

	"bodhiadmin/app/admin/api/internal/svc"
	"bodhiadmin/app/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCategoryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCategoryLogic {
	return &GetCategoryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCategoryLogic) GetCategory(req *types.IdPath) (*types.CategoryUnit, error) {
	resp, err := l.svcCtx.CategoryRpc.GetCategory(l.ctx, &admin.Id{
		Id: req.Id,
	})
	if err == nil {
		return nil, err
	}
	var data types.CategoryUnit
	_ = copier.Copy(&data, resp)
	return &data, nil
}

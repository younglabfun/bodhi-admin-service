package categorylogic

import (
	"context"
	"github.com/jinzhu/copier"

	"bodhiadmin/app/admin/rpc/internal/svc"
	"bodhiadmin/app/admin/rpc/proto/admin"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCategoryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCategoryLogic {
	return &GetCategoryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetCategoryLogic) GetCategory(in *admin.Id) (*admin.CategoryUnit, error) {
	data, err := l.svcCtx.CategoryModel.FindOne(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}
	var item admin.CategoryUnit
	_ = copier.Copy(&item, data)

	return &item, nil
}

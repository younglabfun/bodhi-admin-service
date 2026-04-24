package categorylogic

import (
	"bodhiadmin/app/admin/rpc/model"
	"context"
	"github.com/jinzhu/copier"

	"bodhiadmin/app/admin/rpc/internal/svc"
	"bodhiadmin/app/admin/rpc/proto/admin"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateCategoryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateCategoryLogic {
	return &UpdateCategoryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateCategoryLogic) UpdateCategory(in *admin.CategoryReq) (*admin.AffectedResp, error) {
	var data model.Category
	_ = copier.Copy(&data, in)
	affected := false
	err := l.svcCtx.CategoryModel.Update(l.ctx, &data)
	if err == nil {
		affected = true
	}

	return &admin.AffectedResp{
		Affected: affected,
	}, err
}

package categorylogic

import (
	"bodhiadmin/app/admin/rpc/model"
	"context"
	"github.com/jinzhu/copier"

	"bodhiadmin/app/admin/rpc/internal/svc"
	"bodhiadmin/app/admin/rpc/proto/admin"

	"github.com/zeromicro/go-zero/core/logx"
)

type InsertCategoryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewInsertCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InsertCategoryLogic {
	return &InsertCategoryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *InsertCategoryLogic) InsertCategory(in *admin.CategoryReq) (*admin.AffectedResp, error) {
	var data model.Category
	_ = copier.Copy(&data, in)
	affected := false
	err := l.svcCtx.CategoryModel.Insert(l.ctx, &data)
	if err == nil {
		affected = true
	}

	return &admin.AffectedResp{
		Affected: affected,
	}, err
}

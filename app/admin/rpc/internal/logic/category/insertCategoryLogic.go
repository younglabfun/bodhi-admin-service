package categorylogic

import (
	"bodhiadmin/app/admin/rpc/internal/svc"
	"bodhiadmin/app/admin/rpc/model"
	"bodhiadmin/app/admin/rpc/proto/admin"
	"context"
	"github.com/jinzhu/copier"

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
	affected := false
	//category, err := l.svcCtx.CategoryModel.FindOneByClass(l.ctx, in.Class)
	//if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
	//	return nil, err
	//}
	//if category.Id == in.Id {
	//	err = errors.New("invalid same class")
	//} else {
	var data model.Category
	_ = copier.Copy(&data, in)
	err := l.svcCtx.CategoryModel.Insert(l.ctx, &data)
	if err == nil {
		affected = true
	}
	//}

	return &admin.AffectedResp{
		Affected: affected,
	}, err
}

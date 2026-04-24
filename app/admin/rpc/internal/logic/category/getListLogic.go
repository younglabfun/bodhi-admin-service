package categorylogic

import (
	"context"
	"github.com/jinzhu/copier"

	"bodhiadmin/app/admin/rpc/internal/svc"
	"bodhiadmin/app/admin/rpc/proto/admin"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetListLogic {
	return &GetListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetListLogic) GetList(in *admin.Empty) (*admin.ListCategoryResp, error) {
	resp, err := l.svcCtx.CategoryModel.FindAll(l.ctx)
	if err != nil {
		return nil, err
	}
	var list []*admin.CategoryUnit
	for _, v := range resp {
		var item admin.CategoryUnit
		_ = copier.Copy(&item, v)
		list = append(list, &item)
	}

	return &admin.ListCategoryResp{
		List: list,
	}, nil
}

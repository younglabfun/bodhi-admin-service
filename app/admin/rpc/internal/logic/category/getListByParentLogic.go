package categorylogic

import (
	"context"
	"github.com/jinzhu/copier"

	"bodhiadmin/app/admin/rpc/internal/svc"
	"bodhiadmin/app/admin/rpc/proto/admin"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetListByParentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetListByParentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetListByParentLogic {
	return &GetListByParentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetListByParentLogic) GetListByParent(in *admin.Id) (*admin.ListCategoryResp, error) {
	resp, err := l.svcCtx.CategoryModel.FindListByPid(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}
	var list []*admin.CategoryUnit
	for _, v := range resp {
		var item admin.CategoryUnit
		_ = copier.Copy(&item, v)
		item.HasChildren = false
		if len(v.Children) != 0 {
			item.HasChildren = true
		}
		list = append(list, &item)
	}

	return &admin.ListCategoryResp{
		List: list,
	}, nil
}

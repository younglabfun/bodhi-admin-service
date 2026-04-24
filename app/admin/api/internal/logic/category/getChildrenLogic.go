package category

import (
	"bodhiadmin/app/admin/rpc/proto/admin"
	"context"
	"github.com/jinzhu/copier"

	"bodhiadmin/app/admin/api/internal/svc"
	"bodhiadmin/app/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetChildrenLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetChildrenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetChildrenLogic {
	return &GetChildrenLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetChildrenLogic) GetChildren(req *types.Id) (*types.CategoryListResp, error) {
	resp, err := l.svcCtx.CategoryRpc.GetListByParent(l.ctx, &admin.Id{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}
	var list []*types.CategoryUnit
	for _, v := range resp.List {
		var item types.CategoryUnit
		_ = copier.Copy(&item, v)
		list = append(list, &item)
	}

	return &types.CategoryListResp{
		List: list,
	}, nil
}

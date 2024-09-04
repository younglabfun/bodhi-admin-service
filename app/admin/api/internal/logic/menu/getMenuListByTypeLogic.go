package menu

import (
	"bodhiadmin/app/admin/rpc/proto/admin"
	"bodhiadmin/common/utils"
	"context"
	"github.com/jinzhu/copier"

	"bodhiadmin/app/admin/api/internal/svc"
	"bodhiadmin/app/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMenuListByTypeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetMenuListByTypeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMenuListByTypeLogic {
	return &GetMenuListByTypeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetMenuListByTypeLogic) GetMenuListByType(req *types.ListTypeReq) (*types.MenuTreeResp, error) {
	resp, err := l.svcCtx.MenuRpc.GetMenuListByType(l.ctx, &admin.MenuTypeResp{
		MenuType: req.Type,
	})
	if err != nil {
		return nil, err
	}
	var data []*types.MenuTreeUnit
	if len(resp.List) != 0 {
		for _, v := range resp.List {
			if v.Pid != 0 {
				continue
			}
			var item types.MenuTreeUnit
			_ = copier.Copy(&item, v)
			item.CreatedAt = utils.UnixToStr(v.CreatedAt)

			var children []*types.MenuUnit
			for _, c := range resp.List {
				if c.Pid != v.Id {
					continue
				}
				var child types.MenuUnit
				_ = copier.Copy(&child, c)
				child.CreatedAt = utils.UnixToStr(c.CreatedAt)

				children = append(children, &child)
			}
			item.Children = children

			data = append(data, &item)
		}
	}

	return &types.MenuTreeResp{
		Tree: data,
	}, nil
}

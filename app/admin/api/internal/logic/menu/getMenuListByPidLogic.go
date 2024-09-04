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

type GetMenuListByPidLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetMenuListByPidLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMenuListByPidLogic {
	return &GetMenuListByPidLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetMenuListByPidLogic) GetMenuListByPid(req *types.ListMenuReq) (*types.ListMenuResp, error) {
	resp, err := l.svcCtx.MenuRpc.GetMenuListByPid(l.ctx, &admin.ListMenuReq{
		Pid:      req.Pid,
		MenuType: req.Type,
	})
	if err != nil {
		return nil, err
	}
	var list []*types.MenuUnit
	if len(resp.List) != 0 {
		for _, v := range resp.List {
			var item types.MenuUnit
			_ = copier.Copy(&item, v)
			item.CreatedAt = utils.UnixToStr(v.CreatedAt)

			list = append(list, &item)
		}
	}

	return &types.ListMenuResp{
		List: list,
	}, nil
}

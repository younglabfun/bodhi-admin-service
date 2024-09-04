package menulogic

import (
	"context"
	"github.com/jinzhu/copier"

	"bodhiadmin/app/admin/rpc/internal/svc"
	"bodhiadmin/app/admin/rpc/proto/admin"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMenuListByPidLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetMenuListByPidLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMenuListByPidLogic {
	return &GetMenuListByPidLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetMenuListByPidLogic) GetMenuListByPid(in *admin.ListMenuReq) (*admin.MenuListResp, error) {
	resp, err := l.svcCtx.MenuModel.FindListByPid(l.ctx, in.Pid, in.MenuType)
	if err != nil {
		return nil, err
	}
	var list []*admin.MenuUnit
	if len(resp) != 0 {
		for _, v := range resp {
			var item admin.MenuUnit
			_ = copier.Copy(&item, v)

			list = append(list, &item)
		}
	}

	return &admin.MenuListResp{
		List: list,
	}, nil
}

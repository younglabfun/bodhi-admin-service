package role

import (
	"bodhiadmin/app/admin/rpc/proto/admin"
	"context"
	"github.com/jinzhu/copier"

	"bodhiadmin/app/admin/api/internal/svc"
	"bodhiadmin/app/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetListLogic {
	return &GetListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetListLogic) GetList() (*types.RoleListResp, error) {
	resp, err := l.svcCtx.RoleRpc.GetList(l.ctx, &admin.Empty{})
	if err != nil {
		return nil, err
	}
	var list []*types.RoleUnit
	if len(resp.List) != 0 {
		for _, v := range resp.List {
			var item types.RoleUnit
			_ = copier.Copy(&item, v)

			list = append(list, &item)
		}
	}

	return &types.RoleListResp{
		List: list,
	}, nil
}

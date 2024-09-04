package node

import (
	"bodhiadmin/app/admin/rpc/proto/admin"
	"bodhiadmin/common/utils"
	"context"
	"github.com/jinzhu/copier"

	"bodhiadmin/app/admin/api/internal/svc"
	"bodhiadmin/app/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetListDataLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetListDataLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetListDataLogic {
	return &GetListDataLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetListDataLogic) GetListData() (*types.NodeListResp, error) {
	resp, err := l.svcCtx.NodeRpc.GetList(l.ctx, &admin.Empty{})
	if err != nil {
		return nil, err
	}
	var list []*types.NodeUnit
	if len(resp.List) != 0 {
		for _, v := range resp.List {
			var item types.NodeUnit
			_ = copier.Copy(&item, v)
			item.CreatedAt = utils.UnixToStr(v.CreatedAt)

			list = append(list, &item)
		}
	}

	return &types.NodeListResp{
		List: list,
	}, nil
}

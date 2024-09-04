package nodeGroup

import (
	"bodhiadmin/app/admin/rpc/proto/admin"
	"context"
	"github.com/jinzhu/copier"

	"bodhiadmin/app/admin/api/internal/svc"
	"bodhiadmin/app/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetGroupListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetGroupListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetGroupListLogic {
	return &GetGroupListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetGroupListLogic) GetGroupList() (*types.ListNodeGroupResp, error) {
	resp, err := l.svcCtx.NodeGroupRpc.GetNodeGroupList(l.ctx, &admin.Empty{})
	if err != nil {
		return nil, err
	}

	var list []*types.NodeGroupUnit
	if len(resp.List) != 0 {
		for _, v := range resp.List {
			var item types.NodeGroupUnit
			_ = copier.Copy(&item, v)

			list = append(list, &item)
		}
	}

	return &types.ListNodeGroupResp{
		List: list,
	}, nil
}

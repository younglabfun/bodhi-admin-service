package node

import (
	"bodhiadmin/app/admin/rpc/proto/admin"
	"context"
	"github.com/jinzhu/copier"

	"bodhiadmin/app/admin/api/internal/svc"
	"bodhiadmin/app/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetNodeListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetNodeListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetNodeListLogic {
	return &GetNodeListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetNodeListLogic) GetNodeList(req *types.IdPath) (*types.ListNodeResp, error) {
	resp, err := l.svcCtx.NodeRpc.GetNodeListByGid(l.ctx, &admin.Id{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}

	var list []*types.NodeUnit
	if len(resp.List) != 0 {
		for _, v := range resp.List {
			var item types.NodeUnit
			_ = copier.Copy(&item, v)

			list = append(list, &item)
		}
	}

	return &types.ListNodeResp{
		List: list,
	}, nil
}

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

type ListNodeByGroupLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListNodeByGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListNodeByGroupLogic {
	return &ListNodeByGroupLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListNodeByGroupLogic) ListNodeByGroup(req *types.PageReq) (*types.ListNodeResp, error) {
	var pageReq admin.PageReq
	_ = copier.Copy(&pageReq, req)
	resp, err := l.svcCtx.NodeRpc.ListNode(l.ctx, &pageReq)
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

	return &types.ListNodeResp{
		List:  list,
		Total: resp.Total,
	}, nil
}

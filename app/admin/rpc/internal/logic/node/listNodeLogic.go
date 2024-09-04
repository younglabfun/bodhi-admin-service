package nodelogic

import (
	"bodhiadmin/app/admin/rpc/model"
	"context"
	"github.com/jinzhu/copier"

	"bodhiadmin/app/admin/rpc/internal/svc"
	"bodhiadmin/app/admin/rpc/proto/admin"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListNodeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListNodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListNodeLogic {
	return &ListNodeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ListNodeLogic) ListNode(in *admin.PageReq) (*admin.ListNodeResp, error) {
	var req model.PageReq
	_ = copier.Copy(&req, in)
	resp, total, err := l.svcCtx.NodeModel.FindListByPage(l.ctx, req)
	if err != nil {
		return nil, err
	}
	var list []*admin.NodeUnit
	if len(resp) != 0 {
		for _, v := range resp {
			var item admin.NodeUnit
			_ = copier.Copy(&item, v)

			list = append(list, &item)
		}
	}

	return &admin.ListNodeResp{
		List:  list,
		Total: total,
	}, nil
}

package nodelogic

import (
	"context"
	"github.com/jinzhu/copier"

	"bodhiadmin/app/admin/rpc/internal/svc"
	"bodhiadmin/app/admin/rpc/proto/admin"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetListLogic {
	return &GetListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetListLogic) GetList(in *admin.Empty) (*admin.NodeListResp, error) {
	resp, err := l.svcCtx.NodeModel.FindAll(l.ctx)
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

	return &admin.NodeListResp{
		List: list,
	}, nil
}

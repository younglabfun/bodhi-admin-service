package medialogic

import (
	"bodhiadmin/app/admin/rpc/model"
	"context"
	"github.com/jinzhu/copier"

	"bodhiadmin/app/admin/rpc/internal/svc"
	"bodhiadmin/app/admin/rpc/proto/admin"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListMediaLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListMediaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListMediaLogic {
	return &ListMediaLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ListMediaLogic) ListMedia(in *admin.PageReq) (*admin.ListMediaResp, error) {
	var req model.PageReq
	_ = copier.Copy(&req, in)
	resp, total, err := l.svcCtx.MediaModel.FindListByPage(l.ctx, req)
	if err != nil {
		return nil, err
	}
	var list []*admin.MediaUnit
	if len(resp) != 0 {
		for _, v := range resp {
			var item admin.MediaUnit
			_ = copier.Copy(&item, v)

			list = append(list, &item)
		}
	}

	return &admin.ListMediaResp{
		List:  list,
		Total: total,
	}, nil
}

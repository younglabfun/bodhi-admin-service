package medialogic

import (
	"context"
	"github.com/jinzhu/copier"

	"bodhiadmin/app/admin/rpc/internal/svc"
	"bodhiadmin/app/admin/rpc/proto/admin"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMediaLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetMediaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMediaLogic {
	return &GetMediaLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetMediaLogic) GetMedia(in *admin.Id) (*admin.MediaUnit, error) {
	data, err := l.svcCtx.MediaModel.FindOne(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}
	var item admin.MediaUnit
	_ = copier.Copy(&item, data)

	return &item, nil
}

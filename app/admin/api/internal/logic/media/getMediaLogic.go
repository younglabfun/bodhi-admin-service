package media

import (
	"bodhiadmin/app/admin/rpc/proto/admin"
	"context"
	"github.com/jinzhu/copier"

	"bodhiadmin/app/admin/api/internal/svc"
	"bodhiadmin/app/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMediaLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetMediaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMediaLogic {
	return &GetMediaLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetMediaLogic) GetMedia(req *types.IdPath) (resp *types.MediaUnit, err error) {
	data, err := l.svcCtx.MediaRpc.GetMedia(l.ctx, &admin.Id{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}
	var media types.MediaUnit
	_ = copier.Copy(&media, data)

	return &media, nil
}

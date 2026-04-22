package medialogic

import (
	"bodhiadmin/app/admin/rpc/model"
	"context"
	"github.com/jinzhu/copier"

	"bodhiadmin/app/admin/rpc/internal/svc"
	"bodhiadmin/app/admin/rpc/proto/admin"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateMediaLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateMediaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateMediaLogic {
	return &UpdateMediaLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateMediaLogic) UpdateMedia(in *admin.MediaReq) (*admin.AffectedResp, error) {
	var data model.Media
	_ = copier.Copy(&data, in)
	err := l.svcCtx.MediaModel.Update(l.ctx, &data)

	if err != nil {
		return nil, err
	}

	return &admin.AffectedResp{
		Affected: true,
	}, nil
}

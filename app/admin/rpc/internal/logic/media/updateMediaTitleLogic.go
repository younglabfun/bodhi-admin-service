package medialogic

import (
	"context"

	"bodhiadmin/app/admin/rpc/internal/svc"
	"bodhiadmin/app/admin/rpc/proto/admin"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateMediaTitleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateMediaTitleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateMediaTitleLogic {
	return &UpdateMediaTitleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateMediaTitleLogic) UpdateMediaTitle(in *admin.MediaTitleReq) (*admin.AffectedResp, error) {
	err := l.svcCtx.MediaModel.UpdateTitle(l.ctx, in.Id, in.Title)

	if err != nil {
		return nil, err
	}

	return &admin.AffectedResp{
		Affected: true,
	}, nil
}

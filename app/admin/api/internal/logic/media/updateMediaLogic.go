package media

import (
	"bodhiadmin/app/admin/rpc/proto/admin"
	"context"

	"bodhiadmin/app/admin/api/internal/svc"
	"bodhiadmin/app/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateMediaLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateMediaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateMediaLogic {
	return &UpdateMediaLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateMediaLogic) UpdateMedia(req *types.MediaReq) (*types.AffectedResp, error) {
	resp, err := l.svcCtx.MediaRpc.UpdateMediaTitle(l.ctx, &admin.MediaTitleReq{
		Id:    req.Id,
		Title: req.Title,
	})
	if err != nil {
		return nil, err
	}

	return &types.AffectedResp{
		Affected: resp.Affected,
	}, nil
}

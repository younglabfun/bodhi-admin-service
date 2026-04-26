package article

import (
	"bodhiadmin/app/admin/rpc/proto/admin"
	"context"

	"bodhiadmin/app/admin/api/internal/svc"
	"bodhiadmin/app/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateStatusLogic {
	return &UpdateStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateStatusLogic) UpdateStatus(req *types.StatusReq) (resp *types.AffectedResp, err error) {
	res, err := l.svcCtx.ArticleRpc.UpdateStatus(l.ctx, &admin.StatusReq{
		Id:     req.Id,
		Status: req.Status,
	})
	return &types.AffectedResp{
		Affected: res.Affected,
	}, err
}

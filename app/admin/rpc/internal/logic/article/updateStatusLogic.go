package articlelogic

import (
	"context"

	"bodhiadmin/app/admin/rpc/internal/svc"
	"bodhiadmin/app/admin/rpc/proto/admin"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateStatusLogic {
	return &UpdateStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateStatusLogic) UpdateStatus(in *admin.StatusReq) (*admin.AffectedResp, error) {
	affected := false
	err := l.svcCtx.ArticleModel.UpdateStatus(l.ctx, in.Id, in.Status)
	if err == nil {
		affected = true
	}
	return &admin.AffectedResp{
		Affected: affected,
	}, err
}

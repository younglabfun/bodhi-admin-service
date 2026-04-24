package medialogic

import (
	"context"

	"bodhiadmin/app/admin/rpc/internal/svc"
	"bodhiadmin/app/admin/rpc/proto/admin"

	"github.com/zeromicro/go-zero/core/logx"
)

type RemoveMediaLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRemoveMediaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RemoveMediaLogic {
	return &RemoveMediaLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RemoveMediaLogic) RemoveMedia(in *admin.Id) (*admin.AffectedResp, error) {
	affected := false
	err := l.svcCtx.MediaModel.Delete(l.ctx, in.Id)

	if err == nil {
		affected = true
	}
	return &admin.AffectedResp{
		Affected: affected,
	}, err
}

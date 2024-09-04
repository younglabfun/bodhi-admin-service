package menulogic

import (
	"context"

	"bodhiadmin/app/admin/rpc/internal/svc"
	"bodhiadmin/app/admin/rpc/proto/admin"

	"github.com/zeromicro/go-zero/core/logx"
)

type RemoveMenuLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRemoveMenuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RemoveMenuLogic {
	return &RemoveMenuLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RemoveMenuLogic) RemoveMenu(in *admin.Id) (*admin.AffectedResp, error) {
	err := l.svcCtx.MenuModel.Delete(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}

	return &admin.AffectedResp{
		Affected: true,
	}, nil
}

package userlogic

import (
	"context"
	"github.com/pkg/errors"

	"bodhiadmin/app/admin/rpc/internal/svc"
	"bodhiadmin/app/admin/rpc/proto/admin"

	"github.com/zeromicro/go-zero/core/logx"
)

type RemoveUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRemoveUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RemoveUserLogic {
	return &RemoveUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RemoveUserLogic) RemoveUser(in *admin.Uuid) (*admin.AffectedResp, error) {
	if in.Uuid == l.svcCtx.Config.AdminConf.Master {
		return nil, errors.New("invalid action")
	}

	err := l.svcCtx.UserModel.Delete(l.ctx, in.Uuid)
	if err != nil {
		return nil, err
	}

	return &admin.AffectedResp{
		Affected: true,
	}, nil
}

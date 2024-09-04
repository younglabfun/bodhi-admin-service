package user

import (
	"bodhiadmin/app/admin/rpc/proto/admin"
	"context"

	"bodhiadmin/app/admin/api/internal/svc"
	"bodhiadmin/app/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RemoveUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRemoveUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RemoveUserLogic {
	return &RemoveUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RemoveUserLogic) RemoveUser(req *types.UuidReq) (*types.AffectedResp, error) {
	resp, err := l.svcCtx.UserRpc.RemoveUser(l.ctx, &admin.Uuid{
		Uuid: req.Uuid,
	})
	if err != nil {
		return nil, err
	}

	return &types.AffectedResp{
		Affected: resp.Affected,
	}, nil
}

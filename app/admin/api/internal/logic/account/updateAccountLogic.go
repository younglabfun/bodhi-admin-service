package account

import (
	"bodhiadmin/app/admin/rpc/proto/admin"
	"bodhiadmin/common/utils"
	"context"

	"bodhiadmin/app/admin/api/internal/svc"
	"bodhiadmin/app/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateAccountLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateAccountLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateAccountLogic {
	return &UpdateAccountLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateAccountLogic) UpdateAccount(req *types.AccountReq) (*types.AffectedResp, error) {
	resp, err := l.svcCtx.UserRpc.UpdateUser(l.ctx, &admin.UserReq{
		UserUuid: utils.AnyToStr(l.ctx.Value("UserUuid")),
		Email:    req.Email,
		Name:     req.Name,
		Avatar:   req.Avatar,
	})
	if err != nil {
		return nil, err
	}

	return &types.AffectedResp{
		Affected: resp.Affected,
	}, nil
}

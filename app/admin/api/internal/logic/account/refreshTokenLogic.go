package account

import (
	"bodhiadmin/app/admin/rpc/proto/admin"
	"bodhiadmin/common/utils"
	"context"
	"github.com/jinzhu/copier"

	"bodhiadmin/app/admin/api/internal/svc"
	"bodhiadmin/app/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RefreshTokenLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRefreshTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RefreshTokenLogic {
	return &RefreshTokenLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RefreshTokenLogic) RefreshToken(req *types.TokenReq) (*types.TokenResp, error) {
	resp, err := l.svcCtx.AccountRpc.RefreshToken(l.ctx, &admin.TokenReq{
		UserUuid: utils.AnyToStr(l.ctx.Value("UserUuid")),
		Token:    req.Token,
	})
	if err != nil {
		return nil, err
	}
	var token types.TokenUnit
	_ = copier.Copy(&token, resp.Token)

	return &types.TokenResp{
		Token: token,
	}, nil
}

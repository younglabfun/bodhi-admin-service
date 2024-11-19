package accountlogic

import (
	"bodhiadmin/common/utils"
	"context"
	"github.com/jinzhu/copier"
	"google.golang.org/grpc/status"

	"bodhiadmin/app/admin/rpc/internal/svc"
	"bodhiadmin/app/admin/rpc/proto/admin"

	"github.com/zeromicro/go-zero/core/logx"
)

type RefreshTokenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRefreshTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RefreshTokenLogic {
	return &RefreshTokenLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RefreshTokenLogic) RefreshToken(in *admin.TokenReq) (*admin.TokenResp, error) {
	//check token
	valid := l.svcCtx.TokenModel.CheckRefreshToken(l.ctx, in.UserUuid, in.Token)
	if !valid {
		return nil, status.Errorf(1008, "refresh token is invalid!")
	}
	//build new token
	var jwtAuth utils.JwtAuth
	_ = copier.Copy(&jwtAuth, l.svcCtx.Config.AuthConf)
	tokenData := utils.UserData{
		UserUuid: in.UserUuid,
	}
	jwtToken := utils.BuildToken(jwtAuth, tokenData)
	if jwtToken == nil {
		return nil, status.Errorf(1007, "build new token is failed!")
	}

	//save refresh token
	err := l.svcCtx.TokenModel.SaveRefreshToken(l.ctx, in.UserUuid, jwtToken.RefreshToken, l.svcCtx.Config.AdminConf.RefreshExpired)
	if err != nil {
		return nil, err
	}

	return &admin.TokenResp{
		Token: &admin.TokenUnit{
			AccessToken:  jwtToken.AccessToken,
			RefreshToken: jwtToken.RefreshToken,
			AccessExpire: jwtToken.AccessExpire,
			RefreshAfter: jwtToken.RefreshAfter,
		},
	}, nil
}

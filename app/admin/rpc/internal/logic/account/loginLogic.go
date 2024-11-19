package accountlogic

import (
	"bodhiadmin/app/admin/rpc/internal/svc"
	"bodhiadmin/app/admin/rpc/proto/admin"
	"bodhiadmin/common/utils"
	"context"
	"github.com/jinzhu/copier"
	"google.golang.org/grpc/status"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *admin.LoginReq) (*admin.LoginResp, error) {
	//check user
	user, _ := l.svcCtx.UserModel.FindOneByFields(l.ctx, in.Username)
	if user == nil {
		return nil, status.Errorf(1004, "the user does not exist, please sign up first.")
	}

	if user.IsEnabled == 0 {
		return nil, status.Errorf(1005, "the user is forbidden to sign in, please contact the administrator")
	}

	//check password
	salt := l.svcCtx.Config.AdminConf.Salt
	check := utils.CheckPassword(in.Password, user.Password, salt)
	if !check {
		return nil, status.Errorf(1006, "the password is incorrect!")
	}

	//build token
	var jwtAuth utils.JwtAuth
	_ = copier.Copy(&jwtAuth, l.svcCtx.Config.AuthConf)
	tokenData := utils.UserData{
		UserUuid: user.UserUuid,
	}
	jwtToken := utils.BuildToken(jwtAuth, tokenData)
	if jwtToken == nil {
		return nil, status.Errorf(1007, "build user token is failed!")
	}

	//save refresh token
	err := l.svcCtx.TokenModel.SaveRefreshToken(l.ctx, user.UserUuid, jwtToken.RefreshToken, l.svcCtx.Config.AdminConf.RefreshExpired)
	if err != nil {
		return nil, err
	}
	//save login data
	err = l.svcCtx.UserModel.UpdateLoginData(l.ctx, user.UserUuid, in.ClientIp)
	if err != nil {
		l.Logger.Errorf("update user login data error! ERR: %s", err)
	}

	return &admin.LoginResp{
		User: &admin.UserResp{
			UserUuid: user.UserUuid,
			Username: user.Username,
			Name:     user.Name,
			Avatar:   user.Avatar,
		},
		Token: &admin.TokenUnit{
			AccessToken:  jwtToken.AccessToken,
			RefreshToken: jwtToken.RefreshToken,
			AccessExpire: jwtToken.AccessExpire,
			RefreshAfter: jwtToken.RefreshAfter,
		},
	}, nil
}

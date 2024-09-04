package account

import (
	"bodhiadmin/app/admin/rpc/proto/admin"
	"bodhiadmin/common/utils"
	"context"
	"github.com/jinzhu/copier"
	"net/http"

	"bodhiadmin/app/admin/api/internal/svc"
	"bodhiadmin/app/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq, r *http.Request) (*types.LoginResp, error) {
	clientIp := utils.GetRemoteIp(r)
	//fmt.Println("ip", clientIp)
	resp, err := l.svcCtx.AccountRpc.Login(l.ctx, &admin.LoginReq{
		Username: req.Username,
		Password: req.Password,
		ClientIp: clientIp,
	})
	if err != nil {
		return nil, err
	}

	var loginResp types.LoginResp
	_ = copier.Copy(&loginResp, resp.User)
	_ = copier.Copy(&loginResp.Token, resp.Token)

	return &loginResp, nil
}

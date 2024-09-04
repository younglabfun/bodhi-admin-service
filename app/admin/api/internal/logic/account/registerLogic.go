package account

import (
	"bodhiadmin/app/admin/rpc/proto/admin"
	"bodhiadmin/common/utils"
	"context"
	"google.golang.org/grpc/status"

	"bodhiadmin/app/admin/api/internal/svc"
	"bodhiadmin/app/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterReq) (*types.AffectedResp, error) {
	//check username valid
	if !utils.ValidatorUserName(req.Username, 4, 32) {
		return nil, status.Errorf(1010, "username is invalid!")
	}
	//check email valid
	if !utils.ValidatorEmail(req.Email) {
		return nil, status.Errorf(1012, "email is invalid!")
	}
	//check password valid
	if !utils.ValidatorPassword(req.Password, 8, 16) {
		return nil, status.Errorf(1009, "password is invalid!")
	}
	//check name
	if !utils.ValidatorUserName(req.Name, 4, 32) {
		return nil, status.Errorf(1011, "name is invalid!")
	}

	resp, err := l.svcCtx.AccountRpc.Register(l.ctx, &admin.RegisterReq{
		Username: req.Username,
		Password: req.Password,
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

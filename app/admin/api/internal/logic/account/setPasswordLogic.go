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

type SetPasswordLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSetPasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetPasswordLogic {
	return &SetPasswordLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SetPasswordLogic) SetPassword(req *types.PasswordReq) (*types.AffectedResp, error) {
	//check new password valid
	if !utils.ValidatorPassword(req.NewPassword, 8, 16) {
		return nil, status.Errorf(1009, "the password is invalid!")
	}

	resp, err := l.svcCtx.AccountRpc.SetPassword(l.ctx, &admin.PasswordReq{
		UserUuid:       utils.AnyToStr(l.ctx.Value("UserUuid")),
		Password:       req.Password,
		NewPassword:    req.NewPassword,
		VerifyPassword: true,
	})
	if err != nil {
		return nil, err
	}

	return &types.AffectedResp{
		Affected: resp.Affected,
	}, nil
}

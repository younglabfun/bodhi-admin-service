package accountlogic

import (
	"bodhiadmin/common/utils"
	"context"
	"google.golang.org/grpc/status"

	"bodhiadmin/app/admin/rpc/internal/svc"
	"bodhiadmin/app/admin/rpc/proto/admin"

	"github.com/zeromicro/go-zero/core/logx"
)

type SetPasswordLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSetPasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetPasswordLogic {
	return &SetPasswordLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SetPasswordLogic) SetPassword(in *admin.PasswordReq) (*admin.AffectedResp, error) {

	//check password if VerifyPassword
	if in.VerifyPassword {
		user, err := l.svcCtx.UserModel.FindOne(l.ctx, in.UserUuid)
		if err != nil {
			return nil, err
		}
		valid := utils.CheckPassword(in.Password, user.Password, l.svcCtx.Config.AdminConf.Salt)
		if !valid {
			return nil, status.Errorf(1006, "the password is incorrect!")
		}
	}

	//save new password
	salt := l.svcCtx.Config.AdminConf.Salt
	err := l.svcCtx.UserModel.UpdatePassword(l.ctx, in.UserUuid, in.NewPassword, salt)
	if err != nil {
		return nil, err
	}

	return &admin.AffectedResp{
		Affected: true,
	}, nil
}

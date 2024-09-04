package user

import (
	"bodhiadmin/app/admin/rpc/proto/admin"
	"bodhiadmin/common/utils"
	"context"
	"google.golang.org/grpc/status"

	"bodhiadmin/app/admin/api/internal/svc"
	"bodhiadmin/app/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdatePasswordLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdatePasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdatePasswordLogic {
	return &UpdatePasswordLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdatePasswordLogic) UpdatePassword(req *types.UserPasswordReq) (*types.AffectedResp, error) {
	//check new password valid
	if !utils.ValidatorPassword(req.Password, 8, 16) {
		return nil, status.Errorf(1009, "the password is invalid!")
	}

	resp, err := l.svcCtx.UserRpc.UpdatePassword(l.ctx, &admin.UserPasswordReq{
		UserUuid: req.Uuid,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}

	return &types.AffectedResp{
		Affected: resp.Affected,
	}, nil
}

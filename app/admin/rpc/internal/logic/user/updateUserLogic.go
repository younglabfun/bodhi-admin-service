package userlogic

import (
	"bodhiadmin/app/admin/rpc/model"
	"context"
	"errors"
	"fmt"
	"github.com/SpectatorNan/gorm-zero/gormc"
	"github.com/jinzhu/copier"
	"google.golang.org/grpc/status"

	"bodhiadmin/app/admin/rpc/internal/svc"
	"bodhiadmin/app/admin/rpc/proto/admin"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserLogic {
	return &UpdateUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateUserLogic) UpdateUser(in *admin.UserReq) (*admin.AffectedResp, error) {
	resp, err := l.svcCtx.UserModel.FindOneByUsername(l.ctx, in.Username)
	if err != nil && !errors.Is(err, gormc.ErrNotFound) {
		return nil, err
	}
	if resp != nil && resp.UserUuid != in.UserUuid {
		return nil, status.Errorf(1002, "The username is exist!")
	}
	resp, err = l.svcCtx.UserModel.FindOneByEmail(l.ctx, in.Email)
	if err != nil && !errors.Is(err, gormc.ErrNotFound) {
		return nil, err
	}
	if resp != nil && resp.UserUuid != in.UserUuid {
		return nil, status.Errorf(1003, "The email is exist!")
	}

	var user model.User
	_ = copier.Copy(&user, in)
	fmt.Println("uuuu", user)
	err = l.svcCtx.UserModel.Update(l.ctx, &user)
	if err != nil {
		return nil, err
	}

	return &admin.AffectedResp{
		Affected: true,
	}, nil
}

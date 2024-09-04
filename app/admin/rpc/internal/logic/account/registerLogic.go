package accountlogic

import (
	"bodhiadmin/app/admin/rpc/internal/svc"
	"bodhiadmin/app/admin/rpc/model"
	"bodhiadmin/app/admin/rpc/proto/admin"
	"bodhiadmin/common/utils"
	"context"
	"github.com/pkg/errors"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) Register(in *admin.RegisterReq) (*admin.AffectedResp, error) {
	//check same username
	user, err := l.svcCtx.UserModel.FindOneByUsername(l.ctx, in.Username)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	if user != nil {
		return nil, status.Errorf(1002, "username is exist!")
	}

	email, err := l.svcCtx.UserModel.FindOneByEmail(l.ctx, in.Email)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	if email != nil {
		return nil, status.Errorf(1003, "email is exist!")
	}

	salt := l.svcCtx.Config.AdminConf.Salt
	err = l.svcCtx.UserModel.CreateUser(l.ctx, &model.User{
		UserUuid: utils.CreateUuid(),
		Username: in.Username,
		Password: utils.GetHashedPassword(in.Password, salt),
		Email:    in.Email,
		Name:     in.Name,
		Avatar:   in.Avatar,
	})
	if err != nil {
		return nil, err
	}

	return &admin.AffectedResp{
		Affected: true,
	}, nil
}

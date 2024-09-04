package userlogic

import (
	"bodhiadmin/app/admin/rpc/model"
	"bodhiadmin/common/utils"
	"context"
	"errors"
	"github.com/SpectatorNan/gorm-zero/gormc"
	"google.golang.org/grpc/status"

	"bodhiadmin/app/admin/rpc/internal/svc"
	"bodhiadmin/app/admin/rpc/proto/admin"

	"github.com/zeromicro/go-zero/core/logx"
)

type InsertUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewInsertUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InsertUserLogic {
	return &InsertUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *InsertUserLogic) InsertUser(in *admin.NewUserReq) (*admin.AffectedResp, error) {
	resp, err := l.svcCtx.UserModel.FindOneByUsername(l.ctx, in.Username)
	if err != nil && !errors.Is(err, gormc.ErrNotFound) {
		return nil, err
	}
	if resp != nil {
		return nil, status.Errorf(1002, "The username is exist!")
	}
	resp, err = l.svcCtx.UserModel.FindOneByEmail(l.ctx, in.Email)
	if err != nil && !errors.Is(err, gormc.ErrNotFound) {
		return nil, err
	}
	if resp != nil {
		return nil, status.Errorf(1003, "The email is exist!")
	}

	salt := l.svcCtx.Config.AdminConf.Salt
	user := model.User{
		UserUuid: utils.CreateUuid(),
		Username: in.Username,
		Password: utils.GetHashedPassword(in.Password, salt),
		Email:    in.Email,
		Name:     in.Name,
		Remark:   in.Remark,
	}
	err = l.svcCtx.UserModel.CreateUser(l.ctx, &user)
	if err != nil {
		return nil, err
	}

	return &admin.AffectedResp{
		Affected: true,
	}, nil
}

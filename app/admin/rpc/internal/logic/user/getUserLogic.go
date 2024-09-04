package userlogic

import (
	"context"
	"fmt"
	"github.com/jinzhu/copier"

	"bodhiadmin/app/admin/rpc/internal/svc"
	"bodhiadmin/app/admin/rpc/proto/admin"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLogic {
	return &GetUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserLogic) GetUser(in *admin.Uuid) (*admin.UserUnit, error) {
	resp, err := l.svcCtx.UserModel.FindOne(l.ctx, in.Uuid)
	fmt.Println("rpc resp ", resp, err)
	if err != nil {
		return nil, err
	}
	var user admin.UserUnit
	_ = copier.Copy(&user, resp)

	return &user, nil
}

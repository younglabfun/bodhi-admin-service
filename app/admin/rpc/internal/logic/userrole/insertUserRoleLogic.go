package userrolelogic

import (
	"bodhiadmin/app/admin/rpc/model"
	"context"

	"bodhiadmin/app/admin/rpc/internal/svc"
	"bodhiadmin/app/admin/rpc/proto/admin"

	"github.com/zeromicro/go-zero/core/logx"
)

type InsertUserRoleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewInsertUserRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InsertUserRoleLogic {
	return &InsertUserRoleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *InsertUserRoleLogic) InsertUserRole(in *admin.UserRoleReq) (*admin.AffectedResp, error) {
	err := l.svcCtx.UserRoleModel.Insert(l.ctx, &model.UserRole{
		UserUuid: in.UserUuid,
		RoleUuid: in.RoleUuid,
	})
	if err != nil {
		return nil, err
	}

	return &admin.AffectedResp{
		Affected: true,
	}, nil
}

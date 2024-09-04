package userrolelogic

import (
	"bodhiadmin/app/admin/rpc/model"
	"context"

	"bodhiadmin/app/admin/rpc/internal/svc"
	"bodhiadmin/app/admin/rpc/proto/admin"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserRoleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateUserRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserRoleLogic {
	return &UpdateUserRoleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateUserRoleLogic) UpdateUserRole(in *admin.UpdateUserRoleReq) (*admin.AffectedResp, error) {
	err := l.svcCtx.UserRoleModel.Update(l.ctx, &model.UserRole{
		Id:       in.Id,
		RoleUuid: in.RoleUuid,
	})
	if err != nil {
		return nil, err
	}

	return &admin.AffectedResp{
		Affected: true,
	}, nil
}

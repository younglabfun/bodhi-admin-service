package rolelogic

import (
	"bodhiadmin/app/admin/rpc/model"
	"context"

	"bodhiadmin/app/admin/rpc/internal/svc"
	"bodhiadmin/app/admin/rpc/proto/admin"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateRoleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateRoleLogic {
	return &UpdateRoleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateRoleLogic) UpdateRole(in *admin.RoleReq) (*admin.AffectedResp, error) {
	role := model.Role{
		RoleUuid:      in.RoleUuid,
		Name:          in.Name,
		Description:   in.Description,
		AuthorizeJson: in.AuthorizeJson,
	}
	err := l.svcCtx.RoleModel.Update(l.ctx, &role)
	if err != nil {
		return nil, err
	}

	return &admin.AffectedResp{
		Affected: true,
	}, nil
}

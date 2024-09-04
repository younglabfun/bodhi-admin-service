package rolelogic

import (
	"context"
	"github.com/jinzhu/copier"

	"bodhiadmin/app/admin/rpc/internal/svc"
	"bodhiadmin/app/admin/rpc/proto/admin"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetRoleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRoleLogic {
	return &GetRoleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetRoleLogic) GetRole(in *admin.Uuid) (*admin.RoleUnit, error) {
	resp, err := l.svcCtx.RoleModel.FindOne(l.ctx, in.Uuid)
	if err != nil {
		return nil, err
	}
	var role admin.RoleUnit
	_ = copier.Copy(&role, resp)

	return &role, nil
}

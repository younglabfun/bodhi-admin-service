package rolelogic

import (
	"bodhiadmin/app/admin/rpc/model"
	"bodhiadmin/common/utils"
	"context"

	"bodhiadmin/app/admin/rpc/internal/svc"
	"bodhiadmin/app/admin/rpc/proto/admin"

	"github.com/zeromicro/go-zero/core/logx"
)

type InsertRoleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewInsertRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InsertRoleLogic {
	return &InsertRoleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *InsertRoleLogic) InsertRole(in *admin.RoleReq) (*admin.AffectedResp, error) {
	role := model.Role{
		RoleUuid:      utils.CreateUuid(),
		Name:          in.Name,
		Description:   in.Description,
		AuthorizeJson: in.AuthorizeJson,
	}
	err := l.svcCtx.RoleModel.Insert(l.ctx, &role)
	if err != nil {
		return nil, err
	}

	return &admin.AffectedResp{
		Affected: true,
	}, nil
}

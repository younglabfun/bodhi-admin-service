package role

import (
	"bodhiadmin/app/admin/rpc/proto/admin"
	"context"
	"github.com/jinzhu/copier"

	"bodhiadmin/app/admin/api/internal/svc"
	"bodhiadmin/app/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetRoleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRoleLogic {
	return &GetRoleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetRoleLogic) GetRole(req *types.UuidPath) (*types.RoleResp, error) {
	resp, err := l.svcCtx.RoleRpc.GetRole(l.ctx, &admin.Uuid{
		Uuid: req.Uuid,
	})
	if err != nil {
		return nil, err
	}
	var role types.RoleResp
	_ = copier.Copy(&role, resp)

	return &role, nil
}

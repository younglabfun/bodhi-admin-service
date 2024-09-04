package role

import (
	"bodhiadmin/app/admin/rpc/proto/admin"
	"context"

	"bodhiadmin/app/admin/api/internal/svc"
	"bodhiadmin/app/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type InsertRoleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewInsertRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InsertRoleLogic {
	return &InsertRoleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *InsertRoleLogic) InsertRole(req *types.RoleReq) (*types.AffectedResp, error) {
	role := admin.RoleReq{
		Name:          req.Name,
		Description:   req.Description,
		AuthorizeJson: req.AuthorizeJson,
	}
	resp, err := l.svcCtx.RoleRpc.InsertRole(l.ctx, &role)
	if err != nil {
		return nil, err
	}

	return &types.AffectedResp{
		Affected: resp.Affected,
	}, nil
}

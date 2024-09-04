package account

import (
	"bodhiadmin/app/admin/rpc/proto/admin"
	"bodhiadmin/common/utils"
	"context"
	"github.com/jinzhu/copier"

	"bodhiadmin/app/admin/api/internal/svc"
	"bodhiadmin/app/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPermissionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetPermissionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPermissionLogic {
	return &GetPermissionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetPermissionLogic) GetPermission() (*types.PermissionResp, error) {
	resp, err := l.svcCtx.AccountRpc.GetPermission(l.ctx, &admin.Uuid{
		Uuid: utils.AnyToStr(l.ctx.Value("UserUuid")),
	})
	if err != nil {
		return nil, err
	}

	var permission []string
	_ = copier.Copy(&permission, resp.Permission)

	return &types.PermissionResp{
		Permission: permission,
	}, nil
}

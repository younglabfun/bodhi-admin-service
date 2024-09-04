package rolelogic

import (
	"bodhiadmin/app/admin/rpc/model"
	"context"
	"github.com/jinzhu/copier"

	"bodhiadmin/app/admin/rpc/internal/svc"
	"bodhiadmin/app/admin/rpc/proto/admin"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListRoleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListRoleLogic {
	return &ListRoleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ListRoleLogic) ListRole(in *admin.PageReq) (*admin.ListRoleResp, error) {
	var req model.PageReq
	_ = copier.Copy(&req, in)
	resp, total, err := l.svcCtx.RoleModel.FindListByPage(l.ctx, req)
	if err != nil {
		return nil, err
	}

	var list []*admin.RoleUnit
	if len(resp) != 0 {
		for _, v := range resp {
			var item admin.RoleUnit
			_ = copier.Copy(&item, v)

			list = append(list, &item)
		}
	}

	return &admin.ListRoleResp{
		List:  list,
		Total: total,
	}, nil
}

package role

import (
	"bodhiadmin/app/admin/rpc/proto/admin"
	"bodhiadmin/common/utils"
	"context"
	"github.com/jinzhu/copier"

	"bodhiadmin/app/admin/api/internal/svc"
	"bodhiadmin/app/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListRoleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListRoleLogic {
	return &ListRoleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListRoleLogic) ListRole(req *types.PageReq) (*types.ListRoleResp, error) {
	resp, err := l.svcCtx.RoleRpc.ListRole(l.ctx, &admin.PageReq{
		Page:  req.Page,
		Size:  req.Size,
		Sort:  req.Sort,
		Order: req.Order,
		Field: req.Field,
		Value: req.Value,
	})
	if err != nil {
		return nil, err
	}
	var list []*types.RoleUnit
	if len(resp.List) != 0 {
		for _, v := range resp.List {
			var item types.RoleUnit
			_ = copier.Copy(&item, v)
			item.CreatedAt = utils.UnixToStr(v.CreatedAt)

			list = append(list, &item)
		}
	}

	return &types.ListRoleResp{
		List:  list,
		Total: resp.Total,
	}, nil
}

package user

import (
	"bodhiadmin/app/admin/rpc/proto/admin"
	"bodhiadmin/common/utils"
	"context"
	"github.com/jinzhu/copier"

	"bodhiadmin/app/admin/api/internal/svc"
	"bodhiadmin/app/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListUserLogic {
	return &ListUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListUserLogic) ListUser(req *types.PageReq) (*types.ListUserResp, error) {
	resp, err := l.svcCtx.UserRpc.ListUser(l.ctx, &admin.PageReq{
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
	var list []*types.UserUnit
	if len(resp.List) != 0 {
		for _, v := range resp.List {
			var item types.UserUnit
			_ = copier.Copy(&item, v)
			item.CreatedAt = utils.UnixToStr(v.CreatedAt)

			list = append(list, &item)
		}
	}

	return &types.ListUserResp{
		List:  list,
		Total: resp.Total,
	}, nil
}

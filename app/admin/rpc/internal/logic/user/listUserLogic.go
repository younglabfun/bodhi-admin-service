package userlogic

import (
	"bodhiadmin/app/admin/rpc/model"
	"context"
	"github.com/jinzhu/copier"

	"bodhiadmin/app/admin/rpc/internal/svc"
	"bodhiadmin/app/admin/rpc/proto/admin"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListUserLogic {
	return &ListUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ListUserLogic) ListUser(in *admin.PageReq) (*admin.ListUserResp, error) {
	var req model.PageReq
	_ = copier.Copy(&req, in)
	resp, total, err := l.svcCtx.UserModel.FindListByPage(l.ctx, req)
	if err != nil {
		return nil, err
	}
	var list []*admin.UserUnit
	if len(resp) != 0 {
		for _, v := range resp {
			var item admin.UserUnit
			_ = copier.Copy(&item, v)

			list = append(list, &item)
		}
	}

	return &admin.ListUserResp{
		List:  list,
		Total: total,
	}, nil
}

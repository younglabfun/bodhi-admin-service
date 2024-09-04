package menulogic

import (
	"bodhiadmin/app/admin/rpc/model"
	"context"

	"bodhiadmin/app/admin/rpc/internal/svc"
	"bodhiadmin/app/admin/rpc/proto/admin"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateMenuLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateMenuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateMenuLogic {
	return &UpdateMenuLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateMenuLogic) UpdateMenu(in *admin.MenuReq) (*admin.AffectedResp, error) {
	err := l.svcCtx.MenuModel.Update(l.ctx, &model.Menu{
		Id:        in.Id,
		Pid:       in.Pid,
		Type:      in.Type,
		Title:     in.Title,
		FuncCode:  in.FuncCode,
		Route:     in.Route,
		Component: in.Component,
		Icon:      in.Icon,
		Href:      in.Href,
		Sort:      in.Sort,
		IsShow:    in.IsShow,
	})
	if err != nil {
		return nil, err
	}

	return &admin.AffectedResp{
		Affected: true,
	}, nil
}

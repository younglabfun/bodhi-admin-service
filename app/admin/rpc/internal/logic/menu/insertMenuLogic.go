package menulogic

import (
	"bodhiadmin/app/admin/rpc/model"
	"context"

	"bodhiadmin/app/admin/rpc/internal/svc"
	"bodhiadmin/app/admin/rpc/proto/admin"

	"github.com/zeromicro/go-zero/core/logx"
)

type InsertMenuLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewInsertMenuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InsertMenuLogic {
	return &InsertMenuLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *InsertMenuLogic) InsertMenu(in *admin.MenuReq) (*admin.AffectedResp, error) {
	err := l.svcCtx.MenuModel.Insert(l.ctx, &model.Menu{
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

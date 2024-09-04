package menu

import (
	"bodhiadmin/app/admin/rpc/proto/admin"
	"context"

	"bodhiadmin/app/admin/api/internal/svc"
	"bodhiadmin/app/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type InsertLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewInsertLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InsertLogic {
	return &InsertLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *InsertLogic) Insert(req *types.MenuReq) (*types.AffectedResp, error) {
	resp, err := l.svcCtx.MenuRpc.InsertMenu(l.ctx, &admin.MenuReq{
		Pid:       req.Pid,
		Type:      req.Type,
		Title:     req.Title,
		FuncCode:  req.FuncCode,
		Route:     req.Route,
		Component: req.Component,
		Icon:      req.Icon,
		Href:      req.Href,
		Sort:      req.Sort,
		IsShow:    req.IsShow,
	})
	if err != nil {
		return nil, err
	}

	return &types.AffectedResp{
		Affected: resp.Affected,
	}, nil
}

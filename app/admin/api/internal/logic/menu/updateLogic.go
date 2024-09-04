package menu

import (
	"bodhiadmin/app/admin/rpc/proto/admin"
	"context"

	"bodhiadmin/app/admin/api/internal/svc"
	"bodhiadmin/app/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateLogic {
	return &UpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateLogic) Update(req *types.MenuReq) (*types.AffectedResp, error) {
	resp, err := l.svcCtx.MenuRpc.UpdateMenu(l.ctx, &admin.MenuReq{
		Id:        req.Id,
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

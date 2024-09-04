package nodelogic

import (
	"bodhiadmin/app/admin/rpc/model"
	"context"

	"bodhiadmin/app/admin/rpc/internal/svc"
	"bodhiadmin/app/admin/rpc/proto/admin"

	"github.com/zeromicro/go-zero/core/logx"
)

type InsertNodeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewInsertNodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InsertNodeLogic {
	return &InsertNodeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *InsertNodeLogic) InsertNode(in *admin.NodeReq) (*admin.AffectedResp, error) {
	err := l.svcCtx.NodeModel.Insert(l.ctx, &model.Node{
		GroupId:     in.GroupId,
		FuncCode:    in.FuncCode,
		Name:        in.Name,
		Description: in.Description,
	})
	if err != nil {
		return nil, err
	}

	return &admin.AffectedResp{
		Affected: true,
	}, nil
}

package nodelogic

import (
	"bodhiadmin/app/admin/rpc/model"
	"context"

	"bodhiadmin/app/admin/rpc/internal/svc"
	"bodhiadmin/app/admin/rpc/proto/admin"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateNodeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateNodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateNodeLogic {
	return &UpdateNodeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateNodeLogic) UpdateNode(in *admin.NodeReq) (*admin.AffectedResp, error) {
	err := l.svcCtx.NodeModel.Update(l.ctx, &model.Node{
		Id:          in.Id,
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

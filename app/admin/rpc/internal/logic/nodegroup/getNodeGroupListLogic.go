package nodegrouplogic

import (
	"context"
	"github.com/jinzhu/copier"

	"bodhiadmin/app/admin/rpc/internal/svc"
	"bodhiadmin/app/admin/rpc/proto/admin"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetNodeGroupListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetNodeGroupListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetNodeGroupListLogic {
	return &GetNodeGroupListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetNodeGroupListLogic) GetNodeGroupList(in *admin.Empty) (*admin.NodeGroupList, error) {
	resp, err := l.svcCtx.NodeGroupModel.FindAll(l.ctx)
	if err != nil {
		return nil, err
	}

	var list []*admin.NodeGroupUnit
	if len(resp) != 0 {
		for _, v := range resp {
			var item admin.NodeGroupUnit
			_ = copier.Copy(&item, v)

			list = append(list, &item)
		}
	}

	return &admin.NodeGroupList{
		List: list,
	}, nil
}

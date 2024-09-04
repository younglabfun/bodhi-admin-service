package nodegrouplogic

import (
	"bodhiadmin/app/admin/rpc/internal/svc"
	"bodhiadmin/app/admin/rpc/proto/admin"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type RemoveNodeGroupLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRemoveNodeGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RemoveNodeGroupLogic {
	return &RemoveNodeGroupLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RemoveNodeGroupLogic) RemoveNodeGroup(in *admin.Id) (*admin.AffectedResp, error) {
	//删除之前先将该组数据移至未分组中
	list, err := l.svcCtx.NodeModel.FindListByGid(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}
	if len(list) != 0 {
		var ids []int64
		for _, v := range list {
			ids = append(ids, v.Id)
		}
		err = l.svcCtx.NodeModel.BatchUpdateNodeGroup(l.ctx, ids, 0)
		if err != nil {
			return nil, err
		}
	}

	err = l.svcCtx.NodeGroupModel.Delete(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}

	return &admin.AffectedResp{
		Affected: true,
	}, nil
}

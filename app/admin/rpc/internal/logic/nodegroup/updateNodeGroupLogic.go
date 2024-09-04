package nodegrouplogic

import (
	"bodhiadmin/app/admin/rpc/model"
	"context"
	"errors"
	"gorm.io/gorm"

	"bodhiadmin/app/admin/rpc/internal/svc"
	"bodhiadmin/app/admin/rpc/proto/admin"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateNodeGroupLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateNodeGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateNodeGroupLogic {
	return &UpdateNodeGroupLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateNodeGroupLogic) UpdateNodeGroup(in *admin.NodeGroupReq) (*admin.AffectedResp, error) {
	same, err := l.svcCtx.NodeGroupModel.FindOneByName(l.ctx, in.Name)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	if same != nil && same.Id != in.Id {
		return nil, errors.New("group name is exist")
	}

	group := model.NodeGroup{
		Id:    in.Id,
		Name:  in.Name,
		Title: in.Title,
		Sort:  in.Sort,
	}
	err = l.svcCtx.NodeGroupModel.Update(l.ctx, &group)
	if err != nil {
		return nil, err
	}

	return &admin.AffectedResp{
		Affected: true,
	}, nil
}

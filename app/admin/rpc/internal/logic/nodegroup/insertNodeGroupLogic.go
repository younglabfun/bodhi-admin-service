package nodegrouplogic

import (
	"bodhiadmin/app/admin/rpc/model"
	"context"
	"github.com/pkg/errors"
	"gorm.io/gorm"

	"bodhiadmin/app/admin/rpc/internal/svc"
	"bodhiadmin/app/admin/rpc/proto/admin"

	"github.com/zeromicro/go-zero/core/logx"
)

type InsertNodeGroupLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewInsertNodeGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InsertNodeGroupLogic {
	return &InsertNodeGroupLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *InsertNodeGroupLogic) InsertNodeGroup(in *admin.NodeGroupReq) (*admin.AffectedResp, error) {
	same, err := l.svcCtx.NodeGroupModel.FindOneByName(l.ctx, in.Name)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	if same != nil {
		return nil, errors.New("group name is exist")
	}

	group := model.NodeGroup{
		Name:  in.Name,
		Title: in.Title,
		Sort:  in.Sort,
	}
	err = l.svcCtx.NodeGroupModel.Insert(l.ctx, &group)
	if err != nil {
		return nil, err
	}

	return &admin.AffectedResp{
		Affected: true,
	}, nil
}

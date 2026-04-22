package medialogic

import (
	"bodhiadmin/app/admin/rpc/model"
	"context"
	"github.com/jinzhu/copier"

	"bodhiadmin/app/admin/rpc/internal/svc"
	"bodhiadmin/app/admin/rpc/proto/admin"

	"github.com/zeromicro/go-zero/core/logx"
)

type InsertMediaLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewInsertMediaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InsertMediaLogic {
	return &InsertMediaLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *InsertMediaLogic) InsertMedia(in *admin.MediaReq) (*admin.AffectedResp, error) {
	var data model.Media
	_ = copier.Copy(&data, in)
	err := l.svcCtx.MediaModel.Insert(l.ctx, &data)

	if err != nil {
		return nil, err
	}

	return &admin.AffectedResp{
		Affected: true,
	}, nil
}

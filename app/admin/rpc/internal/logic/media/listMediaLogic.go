package medialogic

import (
	"context"

	"bodhiadmin/app/admin/rpc/internal/svc"
	"bodhiadmin/app/admin/rpc/proto/admin"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListMediaLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListMediaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListMediaLogic {
	return &ListMediaLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ListMediaLogic) ListMedia(in *admin.PageReq) (*admin.ListMediaResp, error) {
	// todo: add your logic here and delete this line

	return &admin.ListMediaResp{}, nil
}

package media

import (
	"context"

	"bodhiadmin/app/admin/api/internal/svc"
	"bodhiadmin/app/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListMediaLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListMediaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListMediaLogic {
	return &ListMediaLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListMediaLogic) ListMedia(req *types.PageReq) (resp *types.ListMediaResp, err error) {
	// todo: add your logic here and delete this line

	return
}

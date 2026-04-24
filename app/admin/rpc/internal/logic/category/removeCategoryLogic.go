package categorylogic

import (
	"context"

	"bodhiadmin/app/admin/rpc/internal/svc"
	"bodhiadmin/app/admin/rpc/proto/admin"

	"github.com/zeromicro/go-zero/core/logx"
)

type RemoveCategoryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRemoveCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RemoveCategoryLogic {
	return &RemoveCategoryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RemoveCategoryLogic) RemoveCategory(in *admin.Id) (*admin.AffectedResp, error) {
	affected := false
	err := l.svcCtx.CategoryModel.Delete(l.ctx, in.Id)
	if err == nil {
		affected = true
	}
	return &admin.AffectedResp{
		Affected: affected,
	}, err
}

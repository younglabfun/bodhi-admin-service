package articlelogic

import (
	"context"

	"bodhiadmin/app/admin/rpc/internal/svc"
	"bodhiadmin/app/admin/rpc/proto/admin"

	"github.com/zeromicro/go-zero/core/logx"
)

type RemoveArticleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRemoveArticleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RemoveArticleLogic {
	return &RemoveArticleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RemoveArticleLogic) RemoveArticle(in *admin.Id) (*admin.AffectedResp, error) {
	affected := false
	err := l.svcCtx.ArticleModel.Delete(l.ctx, in.Id)
	if err == nil {
		affected = true
	}
	return &admin.AffectedResp{
		Affected: affected,
	}, err
}

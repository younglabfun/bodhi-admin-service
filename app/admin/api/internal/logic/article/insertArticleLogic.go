package article

import (
	"bodhiadmin/app/admin/rpc/proto/admin"
	"context"
	"github.com/jinzhu/copier"

	"bodhiadmin/app/admin/api/internal/svc"
	"bodhiadmin/app/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type InsertArticleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewInsertArticleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InsertArticleLogic {
	return &InsertArticleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *InsertArticleLogic) InsertArticle(req *types.ArticleReq) (resp *types.AffectedResp, err error) {
	var article admin.ArticleUnit
	_ = copier.Copy(&article, req)
	res, err := l.svcCtx.ArticleRpc.InsertArticle(l.ctx, &article)

	return &types.AffectedResp{
		Affected: res.Affected,
	}, err
}

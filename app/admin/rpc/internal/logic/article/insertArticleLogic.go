package articlelogic

import (
	"bodhiadmin/app/admin/rpc/model"
	"context"
	"github.com/jinzhu/copier"

	"bodhiadmin/app/admin/rpc/internal/svc"
	"bodhiadmin/app/admin/rpc/proto/admin"

	"github.com/zeromicro/go-zero/core/logx"
)

type InsertArticleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewInsertArticleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InsertArticleLogic {
	return &InsertArticleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *InsertArticleLogic) InsertArticle(in *admin.ArticleUnit) (*admin.AffectedResp, error) {
	affected := false
	var data model.Article
	_ = copier.Copy(&data, in)
	err := l.svcCtx.ArticleModel.InsertWithRelationData(l.ctx, &data, in.CategoryIds)
	if err == nil {
		affected = true
	}

	return &admin.AffectedResp{
		Affected: affected,
	}, err
}

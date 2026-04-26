package articlelogic

import (
	"context"

	"bodhiadmin/app/admin/rpc/internal/svc"
	"bodhiadmin/app/admin/rpc/proto/admin"
	"github.com/fatih/structs"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateArticleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateArticleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateArticleLogic {
	return &UpdateArticleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateArticleLogic) UpdateArticle(in *admin.ArticleUnit) (*admin.AffectedResp, error) {
	article := structs.Map(in)
	delete(article, "CategoryIds")
	err := l.svcCtx.ArticleModel.UpdateArticle(l.ctx, in.Id, article)
	if err != nil {
		return nil, err
	}
	err = l.svcCtx.ArticleCategoryLinkModel.UpdateCategoryLink(l.ctx, in.Id, in.CategoryIds)
	if err != nil {
		return nil, err
	}

	return &admin.AffectedResp{
		Affected: true,
	}, nil
}

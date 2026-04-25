package articlelogic

import (
	"context"
	"github.com/jinzhu/copier"

	"bodhiadmin/app/admin/rpc/internal/svc"
	"bodhiadmin/app/admin/rpc/proto/admin"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetArticleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetArticleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetArticleLogic {
	return &GetArticleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetArticleLogic) GetArticle(in *admin.Id) (*admin.ArticleUnit, error) {
	data, err := l.svcCtx.ArticleModel.FindOne(l.ctx, in.Id, true)
	if err != nil {
		return nil, err
	}
	var item admin.ArticleUnit
	_ = copier.Copy(&item, data)

	return &item, nil
}

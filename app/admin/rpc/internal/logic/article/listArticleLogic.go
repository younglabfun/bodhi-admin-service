package articlelogic

import (
	"bodhiadmin/app/admin/rpc/model"
	"context"
	"github.com/jinzhu/copier"

	"bodhiadmin/app/admin/rpc/internal/svc"
	"bodhiadmin/app/admin/rpc/proto/admin"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListArticleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListArticleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListArticleLogic {
	return &ListArticleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ListArticleLogic) ListArticle(in *admin.PageReq) (*admin.ListArticleResp, error) {
	var req model.PageReq
	_ = copier.Copy(&req, in)

	resp, total, err := l.svcCtx.ArticleModel.FindListByPage(l.ctx, req)
	if err != nil {
		return nil, err
	}
	var list []*admin.ArticleDataUnit
	for _, v := range resp {
		var item admin.ArticleDataUnit
		_ = copier.Copy(&item, v)
		list = append(list, &item)
	}

	return &admin.ListArticleResp{
		List:  list,
		Total: total,
	}, nil
}

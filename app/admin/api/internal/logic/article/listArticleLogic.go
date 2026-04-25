package article

import (
	"bodhiadmin/app/admin/rpc/proto/admin"
	"bodhiadmin/common/utils"
	"context"
	"github.com/jinzhu/copier"

	"bodhiadmin/app/admin/api/internal/svc"
	"bodhiadmin/app/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListArticleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListArticleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListArticleLogic {
	return &ListArticleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListArticleLogic) ListArticle(req *types.PageReq) (*types.ListArticleResp, error) {
	var pageReq admin.PageReq
	_ = copier.Copy(&pageReq, req)
	resp, err := l.svcCtx.ArticleRpc.ListArticle(l.ctx, &pageReq)
	if err != nil {
		return nil, err
	}
	var list []*types.ArticleUnit
	for _, v := range resp.List {
		var item types.ArticleUnit
		_ = copier.Copy(&item, v)
		var categories []*types.ArticleCategory
		if len(v.Categories) != 0 {
			for _, c := range v.Categories {
				var category types.ArticleCategory
				_ = copier.Copy(&category, c)
				categories = append(categories, &category)
			}
		}
		item.Categories = categories
		//fmt.Println("===", v.Categories)
		item.CreatedAt = utils.UnixToStr(v.CreatedAt)

		list = append(list, &item)
	}
	return &types.ListArticleResp{
		List:  list,
		Total: resp.Total,
	}, nil
}

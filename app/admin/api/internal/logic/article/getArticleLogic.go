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

type GetArticleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetArticleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetArticleLogic {
	return &GetArticleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetArticleLogic) GetArticle(req *types.IdPath) (resp *types.ArticleResp, err error) {
	res, err := l.svcCtx.ArticleRpc.GetArticle(l.ctx, &admin.Id{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}
	var article types.ArticleResp
	_ = copier.Copy(&article, res)
	var thumbnails []string
	for k, _ := range utils.ThumbnailConfig {
		pic := utils.GetThumbnailFilename(res.CoverImg, k)
		thumbnails = append(thumbnails, pic)
	}
	article.Thumbnails = thumbnails

	return &article, nil
}

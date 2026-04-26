package article

import (
	"bodhiadmin/app/admin/rpc/proto/admin"
	"context"

	"bodhiadmin/app/admin/api/internal/svc"
	"bodhiadmin/app/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RemoveArticleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRemoveArticleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RemoveArticleLogic {
	return &RemoveArticleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RemoveArticleLogic) RemoveArticle(req *types.IdReq) (*types.AffectedResp, error) {
	resp, err := l.svcCtx.ArticleRpc.RemoveArticle(l.ctx, &admin.Id{
		Id: req.Id,
	})

	return &types.AffectedResp{
		Affected: resp.Affected,
	}, err
}

package media

import (
	"bodhiadmin/app/admin/rpc/proto/admin"
	"bodhiadmin/common/utils"
	"context"
	"github.com/jinzhu/copier"

	"bodhiadmin/app/admin/api/internal/svc"
	"bodhiadmin/app/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListMediaLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListMediaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListMediaLogic {
	return &ListMediaLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListMediaLogic) ListMedia(req *types.PageReq) (*types.ListMediaResp, error) {
	resp, err := l.svcCtx.MediaRpc.ListMedia(l.ctx, &admin.PageReq{
		Page:  req.Page,
		Size:  req.Size,
		Sort:  req.Sort,
		Order: req.Order,
		Field: req.Field,
		Value: req.Value,
	})
	if err != nil {
		return nil, err
	}
	var list []*types.MediaUnit
	if len(resp.List) != 0 {
		for _, v := range resp.List {
			var item types.MediaUnit
			_ = copier.Copy(&item, v)
			small := utils.GetThumbnailFilename(v.Path, "small")
			item.Mini = l.svcCtx.Config.AdminConf.ImagePath + small
			medium := utils.GetThumbnailFilename(v.Path, "medium")
			item.Preview = l.svcCtx.Config.AdminConf.ImagePath + medium
			item.Path = l.svcCtx.Config.AdminConf.ImagePath + v.Path
			item.CreatedAt = utils.UnixToStr(v.CreatedAt)

			list = append(list, &item)
		}
	}

	return &types.ListMediaResp{
		List:  list,
		Total: resp.Total,
	}, nil
}

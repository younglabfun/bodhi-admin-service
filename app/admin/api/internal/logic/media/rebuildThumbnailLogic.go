package media

import (
	"bodhiadmin/app/admin/api/internal/svc"
	"bodhiadmin/app/admin/api/internal/types"
	"bodhiadmin/app/admin/rpc/proto/admin"
	"bodhiadmin/common/utils"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
)

type RebuildThumbnailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRebuildThumbnailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RebuildThumbnailLogic {
	return &RebuildThumbnailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RebuildThumbnailLogic) RebuildThumbnail(req *types.IdPath) (resp *types.AffectedResp, err error) {
	data, err := l.svcCtx.MediaRpc.GetMedia(l.ctx, &admin.Id{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}
	affected := false
	//fmt.Println("get filename ", filename)
	err = utils.DoProcessImage(data.Path, l.svcCtx.Config.AdminConf.UploadPath)
	if err == nil {
		affected = true
	}

	return &types.AffectedResp{
		Affected: affected,
	}, err
}

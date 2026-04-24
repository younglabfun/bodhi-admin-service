package media

import (
	"bodhiadmin/app/admin/rpc/proto/admin"
	"bodhiadmin/common/utils"
	"context"

	"bodhiadmin/app/admin/api/internal/svc"
	"bodhiadmin/app/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RemoveMediaLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRemoveMediaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RemoveMediaLogic {
	return &RemoveMediaLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RemoveMediaLogic) RemoveMedia(req *types.IdReq) (*types.AffectedResp, error) {
	data, err := l.svcCtx.MediaRpc.GetMedia(l.ctx, &admin.Id{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}
	affected := false
	err = utils.RemoveFileAndThumbnails(data.Path, l.svcCtx.Config.AdminConf.UploadPath)
	if err == nil {
		resp, err := l.svcCtx.MediaRpc.RemoveMedia(l.ctx, &admin.Id{
			Id: req.Id,
		})
		if err != nil {
			return nil, err
		}
		affected = resp.Affected
	}

	return &types.AffectedResp{
		Affected: affected,
	}, err
}

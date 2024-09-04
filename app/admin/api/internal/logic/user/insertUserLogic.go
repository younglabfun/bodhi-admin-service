package user

import (
	"bodhiadmin/app/admin/rpc/proto/admin"
	"context"

	"bodhiadmin/app/admin/api/internal/svc"
	"bodhiadmin/app/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type InsertUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewInsertUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InsertUserLogic {
	return &InsertUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *InsertUserLogic) InsertUser(req *types.NewUserReq) (*types.AffectedResp, error) {
	user := admin.NewUserReq{
		Username: req.Username,
		Password: req.Password,
		Email:    req.Email,
		Name:     req.Name,
		Remark:   req.Remark,
	}
	resp, err := l.svcCtx.UserRpc.InsertUser(l.ctx, &user)
	if err != nil {
		return nil, err
	}

	return &types.AffectedResp{
		Affected: resp.Affected,
	}, nil
}

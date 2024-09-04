package account

import (
	"bodhiadmin/app/admin/rpc/proto/admin"
	"bodhiadmin/common/utils"
	"context"
	"github.com/jinzhu/copier"

	"bodhiadmin/app/admin/api/internal/svc"
	"bodhiadmin/app/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAccountLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetAccountLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAccountLogic {
	return &GetAccountLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAccountLogic) GetAccount() (*types.AccountResp, error) {
	//fmt.Println("uuid ", utils.AnyToStr(l.ctx.Value("UserUuid")))
	resp, err := l.svcCtx.UserRpc.GetUser(l.ctx, &admin.Uuid{
		Uuid: utils.AnyToStr(l.ctx.Value("UserUuid")),
	})
	//fmt.Println("resp ", resp, err)
	if err != nil {
		return nil, err
	}
	var account types.AccountResp
	_ = copier.Copy(&account, resp)
	account.Uuid = resp.UserUuid

	return &account, nil
}

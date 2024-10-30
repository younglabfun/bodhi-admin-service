package menu

import (
	"bodhiadmin/app/admin/rpc/proto/admin"
	"bodhiadmin/common/utils"
	"context"
	"github.com/jinzhu/copier"

	"bodhiadmin/app/admin/api/internal/svc"
	"bodhiadmin/app/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAdminMenuLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetAdminMenuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAdminMenuLogic {
	return &GetAdminMenuLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAdminMenuLogic) GetAdminMenu() (*types.MenuTreeResp, error) {
	//fmt.Println("id path", req)
	resp, err := l.svcCtx.MenuRpc.GetMenuListByType(l.ctx, &admin.MenuTypeResp{
		MenuType: 0,
	})
	if err != nil {
		return nil, err
	}
	//get user admin.Permission
	userUuid := utils.AnyToStr(l.ctx.Value("UserUuid"))
	respData, err := l.svcCtx.AccountRpc.GetPermission(l.ctx, &admin.Uuid{
		Uuid: userUuid,
	})
	if err != nil {
		return nil, err
	}
	var permissions = respData.Permission
	//fmt.Println("permi ", permissions)
	//fmt.Println("menu ", len(resp.List))

	var data []*types.MenuTreeUnit
	if len(resp.List) != 0 {
		for _, v := range resp.List {
			if v.IsEnabled != 1 || v.Pid != 0 {
				continue
			}

			//fmt.Println("fun code 111 ", v.FuncCode)
			var item types.MenuTreeUnit
			_ = copier.Copy(&item, v)
			item.CreatedAt = utils.UnixToStr(v.CreatedAt)

			var children []*types.MenuUnit
			for _, c := range resp.List {
				if c.IsEnabled != 1 || c.Pid != v.Id {
					continue
				}

				if permissions == nil || !utils.ContainsStr(permissions, c.FuncCode) {
					continue
				}

				//fmt.Println("fun code ", c.FuncCode)

				var child types.MenuUnit
				_ = copier.Copy(&child, c)
				child.CreatedAt = utils.UnixToStr(c.CreatedAt)

				children = append(children, &child)
			}
			if len(children) == 0 {
				//一级目录，check permissions
				if permissions == nil || !utils.ContainsStr(permissions, v.FuncCode) {
					continue
				}
			} else {
				item.Children = children
			}

			data = append(data, &item)
		}
	}

	return &types.MenuTreeResp{
		Tree: data,
	}, nil
}

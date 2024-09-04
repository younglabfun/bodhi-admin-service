// Code generated by goctl. DO NOT EDIT.
package types

type LoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResp struct {
	UserUuid string    `json:"userUuid"`
	Username string    `json:"username"`
	Name     string    `json:"name"`
	Avatar   string    `json:"avatar"`
	Token    TokenUnit `json:"token"`
}

type TokenUnit struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
	AccessExpire int64  `json:"accessExpire"`
	RefreshAfter int64  `json:"refreshAfter"`
}

type RegisterReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	Avatar   string `json:"avatar"`
}

type TokenReq struct {
	Token string `json:"token"`
}

type TokenResp struct {
	Token TokenUnit `json:"token"`
}

type PasswordReq struct {
	Password    string `json:"password"`
	NewPassword string `json:"newPassword"`
}

type AccountReq struct {
	UserUuid string `json:"userUuid"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	Avatar   string `json:"Avatar"`
}

type AccountResp struct {
	Uuid   string `json:"uuid"`
	Name   string `json:"name"`
	Avatar string `json:"Avatar"`
}

type PermissionResp struct {
	Permission []string `json:"permission"`
}

type PageReq struct {
	Page  int64  `form:"page"`
	Size  int64  `form:"size"`
	Sort  string `form:"sort,optional"`
	Order string `form:"order,optional"`
	Field string `form:"field,optional"`
	Value string `form:"value,optional"`
}

type PageResp struct {
	Total int64 `json:"total"`
}

type AffectedResp struct {
	Affected bool `json:"affected"`
}

type UuidReq struct {
	Uuid string `json:"uuid,optional"`
}

type UuidPath struct {
	Uuid string `path:"uuid"`
}

type Id struct {
	Id int64 `form:"id,optional"`
}

type IdReq struct {
	Id int64 `json:"id"`
}

type IdPath struct {
	Id int64 `path:"id,optional"`
}

type StatusReq struct {
	Id     int64  `json:"id"`
	Status string `json:"status"`
}

type UuidStatusReq struct {
	Uuid   string `json:"uuid"`
	Status string `json:"status"`
}

type MenuReq struct {
	Id        int64  `json:"id,optional"`
	Pid       int64  `json:"pid"`
	Type      int64  `json:"type"`
	Title     string `json:"title"`
	FuncCode  string `json:"funcCode"`
	Route     string `json:"route"`
	Component string `json:"component"`
	Icon      string `json:"icon,optional"`
	Href      string `json:"href,optional"`
	Sort      int64  `json:"sort"`
	IsShow    int64  `json:"isShow"`
}

type MenuResp struct {
	Id        int64  `json:"id"`
	Pid       int64  `json:"pid"`
	Type      int64  `json:"type"`
	Title     string `json:"title"`
	FuncCode  string `json:"funcCode"`
	Route     string `json:"route"`
	Component string `json:"component"`
	Icon      string `json:"icon"`
	Href      string `json:"href"`
	Sort      int64  `json:"sort"`
	IsShow    int64  `json:"isShow"`
}

type MenuUnit struct {
	Id        int64  `json:"id"`
	Pid       int64  `json:"pid"`
	Type      int64  `json:"type"`
	Title     string `json:"title"`
	FuncCode  string `json:"funcCode"`
	Route     string `json:"route"`
	Component string `json:"component"`
	Icon      string `json:"icon"`
	Href      string `json:"href"`
	Sort      int64  `json:"sort"`
	IsShow    int64  `json:"isShow"`
	IsEnabled int64  `json:"isEnabled"`
	CreatedAt string `json:"createdAt"`
}

type ListMenuReq struct {
	Pid  int64 `form:"pid"`
	Type int64 `form:"type"`
}

type ListMenuResp struct {
	List []*MenuUnit `json:"list"`
}

type ListTypeReq struct {
	Type int64 `form:"type"`
}

type MenuTreeUnit struct {
	MenuUnit
	Children []*MenuUnit `json:"children"`
}

type MenuTreeResp struct {
	Tree []*MenuTreeUnit `json:"tree"`
}

type NodeReq struct {
	Id          int64  `json:"id,optional"`
	GroupId     int64  `json:"groupId"`
	FuncCode    string `json:"funcCode"`
	Name        string `json:"name"`
	Description string `json:"description,optional"`
}

type NodeResp struct {
	Id          int64  `json:"id"`
	GroupId     int64  `json:"groupId"`
	FuncCode    string `json:"funcCode"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type NodeUnit struct {
	Id          int64  `json:"id"`
	GroupId     int64  `json:"groupId"`
	FuncCode    string `json:"funcCode"`
	Name        string `json:"name"`
	Description string `json:"description"`
	IsEnabled   int64  `json:"isEnabled"`
	CreatedAt   string `json:"createdAt"`
}

type NodeListResp struct {
	List []*NodeUnit `json:"list"`
}

type ListNodeResp struct {
	List  []*NodeUnit `json:"list"`
	Total int64       `json:"total"`
}

type MoveReq struct {
	GroupId int64  `json:"groupId"`
	Ids     string `json:"ids"`
}

type BatchRemoveReq struct {
	Ids string `json:"ids"`
}

type NodeGroupReq struct {
	Id    int64  `json:"id,optional"`
	Title string `json:"title"`
	Name  string `json:"name"`
	Sort  int64  `json:"sort"`
}

type NodeGroupResp struct {
	Id    int64  `json:"id"`
	Title string `json:"title"`
	Name  string `json:"name"`
	Sort  int64  `json:"sort"`
}

type NodeGroupUnit struct {
	Id    int64  `json:"id"`
	Title string `json:"title"`
	Name  string `json:"name"`
}

type ListNodeGroupResp struct {
	List []*NodeGroupUnit `json:"list"`
}

type RoleReq struct {
	RoleUuid      string `json:"roleUuid,optional"`
	Name          string `json:"name"`
	Description   string `json:"description,optional"`
	AuthorizeJson string `json:"authorizeJson"`
	IsDefault     int64  `json:"isDefault,optional"`
}

type RoleResp struct {
	RoleUuid      string `json:"roleUuid"`
	Name          string `json:"name"`
	Description   string `json:"description"`
	AuthorizeJson string `json:"authorizeJson"`
	IsDefault     int64  `json:"isDefault"`
}

type RoleUnit struct {
	RoleUuid      string `json:"roleUuid,optional"`
	Name          string `json:"name"`
	Description   string `json:"description"`
	AuthorizeJson string `json:"authorizeJson"`
	IsDefault     int64  `json:"isDefault"`
	IsEnabled     int64  `json:"isEnabled"`
	CreatedAt     string `json:"createdAt"`
}

type ListRoleResp struct {
	List  []*RoleUnit `json:"list"`
	Total int64       `json:"total"`
}

type RoleListResp struct {
	List []*RoleUnit `json:"list"`
}

type NewUserReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	Remark   string `json:"remark,optional"`
}

type UserReq struct {
	UserUuid     string `json:"userUuid"`
	Username     string `json:"username"`
	Email        string `json:"email"`
	Name         string `json:"name"`
	Avatar       string `json:"avatar,optional"`
	Remark       string `json:"remark,optional"`
	MailVerified int64  `json:"mailVerified,optional"`
	IsEnabled    int64  `json:"isEnabled,optional"`
}

type UserResp struct {
	UserUuid string `json:"userUuid"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	Avatar   string `json:"avatar"`
	Remark   string `json:"remark"`
}

type UserUnit struct {
	UserUuid  string `json:"userUuid"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Name      string `json:"name"`
	Remark    string `json:"remark"`
	IsEnabled int64  `json:"isEnabled"`
	CreatedAt string `json:"createdAt"`
}

type ListUserResp struct {
	List  []*UserUnit `json:"list"`
	Total int64       `json:"total"`
}

type UserRoleReq struct {
	UserUuid string `json:"userUuid"`
	RoleUuid string `json:"roleUuid"`
}

type ResetRoleReq struct {
	Id       int64  `json:"id"`
	RoleUuid string `json:"roleUuid"`
}

type UserPasswordReq struct {
	Uuid     string `json:"uuid"`
	Password string `json:"password"`
}

type UserRoleUnit struct {
	Id       int64  `json:"id"`
	RoleUuid string `json:"roleUuid"`
}

type UserRolesResp struct {
	List []*UserRoleUnit `json:"list"`
}

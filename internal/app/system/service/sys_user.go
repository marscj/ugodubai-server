package service

import (
	"context"

	"ugodubai-server/api/v1/system"
	"ugodubai-server/internal/app/system/model"
	"ugodubai-server/internal/app/system/model/entity"

	"github.com/gogf/gf/v2/container/gset"
	"github.com/gogf/gf/v2/database/gdb"
)

type (
	ISysUser interface {
		GetCasBinUserPrefix() string
		NotCheckAuthAdminIds(ctx context.Context) *gset.Set
		GetAdminUserByUsernamePassword(ctx context.Context, req *system.UserLoginReq) (user *model.LoginUser, err error)
		GetUserByUsername(ctx context.Context, userName string) (user *model.LoginUser, err error)
		GetUserById(ctx context.Context, id uint64) (user *model.LoginUser, err error)
		LoginLog(ctx context.Context, params *model.LoginLogParams)
		UpdateLoginInfo(ctx context.Context, id uint64, ip string) (err error)
		GetAdminRules(ctx context.Context, userId uint64) (menuList []*model.UserMenus, permissions []string, err error)
		GetAdminRole(ctx context.Context, userId uint64, allRoleList []*entity.SysRole) (roles []*entity.SysRole, err error)
		GetAdminRoleIds(ctx context.Context, userId uint64) (roleIds []uint, err error)
		GetAllMenus(ctx context.Context) (menus []*model.UserMenus, err error)
		GetAdminMenusByRoleIds(ctx context.Context, roleIds []uint) (menus []*model.UserMenus, err error)
		GetMenusTree(menus []*model.UserMenus, pid uint) []*model.UserMenus
		GetPermissions(ctx context.Context, roleIds []uint) (userButtons []string, err error)
		List(ctx context.Context, req *system.UserListReq) (total interface{}, userList []*model.SysUser, err error)
		Add(ctx context.Context, req *system.UserAddReq) (err error)
		Edit(ctx context.Context, req *system.UserEditReq) (err error)
		AddUserPost(ctx context.Context, tx gdb.TX, postIds []int64, userId int64) (err error)
		EditUserRole(ctx context.Context, roleIds []int64, userId int64) (err error)
		UserNameOrMobileExists(ctx context.Context, userName, mobile string, id ...int64) error
		GetEditUser(ctx context.Context, id uint64) (res *system.UserGetEditRes, err error)
		GetUserInfoById(ctx context.Context, id uint64, withPwd ...bool) (user *entity.SysUser, err error)
		GetUserPostIds(ctx context.Context, userId uint64) (postIds []int64, err error)
		ResetUserPwd(ctx context.Context, req *system.UserResetPwdReq) (err error)
		ChangeUserStatus(ctx context.Context, req *system.UserStatusReq) (err error)
		Delete(ctx context.Context, ids []int) (err error)
		GetUsers(ctx context.Context, ids []int) (users []*model.SysUserSimple, err error)
		GetUsersByAgenId(ctx context.Context, id uint64) (users []*model.SysUserSimple, err error)
	}
)

var (
	localSysUser ISysUser
)

func SysUser() ISysUser {
	if localSysUser == nil {
		panic("implement not found for interface ISysUser, forgot register?")
	}
	return localSysUser
}

func RegisterSysUser(i ISysUser) {
	localSysUser = i
}

package system

import (
	commonApi "ugodubai-server/api/v1/common"
	"ugodubai-server/internal/app/system/model/entity"

	"github.com/gogf/gf/v2/frame/g"
)

// SysUserOnlineListReq 列表搜索参数
type SysUserOnlineListReq struct {
	g.Meta   `path:"/online/list" tags:"在线用户管理" method:"get" summary:"列表"`
	Username string `p:"userName"`
	Ip       string `p:"ipaddr"`
	commonApi.PageReq
	commonApi.Author
}

// SysUserOnlineListRes 列表结果
type SysUserOnlineListRes struct {
	g.Meta `mime:"application/json"`
	commonApi.ListRes
	List []*entity.SysUserOnline `json:"list"`
}

type SysUserOnlineForceLogoutReq struct {
	g.Meta `path:"/online/forceLogout" tags:"在线用户管理" method:"delete" summary:"强制用户退出登录"`
	commonApi.Author
	Ids []int `p:"ids" v:"required#ids不能为空"`
}

type SysUserOnlineForceLogoutRes struct {
	commonApi.EmptyRes
}

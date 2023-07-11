package system

import (
	commonApi "ugodubai-server/api/v1/common"
	"ugodubai-server/internal/app/system/model"

	"github.com/gogf/gf/v2/frame/g"
)

// SysOperLogListReq 分页请求参数
type SysOperLogListReq struct {
	g.Meta        `path:"/operLog/list" tags:"操作日志" method:"get" summary:"操作日志列表"`
	Title         string `p:"title"`         //系统模块
	RequestMethod string `p:"requestMethod"` //请求方式
	OperName      string `p:"operName"`      //操作人员
	commonApi.PageReq
	commonApi.Author
}

// SysOperLogListRes 列表返回结果
type SysOperLogListRes struct {
	g.Meta `mime:"application/json"`
	commonApi.ListRes
	List []*model.SysOperLogListModel `json:"list"`
}

// SysOperLogGetReq 获取一条数据请求
type SysOperLogGetReq struct {
	g.Meta `path:"/operLog/get" tags:"操作日志" method:"get" summary:"获取操作日志信息"`
	commonApi.Author
	OperId uint64 `p:"operId" v:"required#主键必须"` //通过主键获取
}

// SysOperLogGetRes 获取一条数据结果
type SysOperLogGetRes struct {
	g.Meta `mime:"application/json"`
	*model.SysOperLogInfoModel
}

// SysOperLogDeleteReq 删除数据请求
type SysOperLogDeleteReq struct {
	g.Meta `path:"/operLog/delete" tags:"操作日志" method:"delete" summary:"删除操作日志"`
	commonApi.Author
	OperIds []uint64 `p:"operIds" v:"required#主键必须"` //通过主键删除
}

// SysOperLogDeleteRes 删除数据返回
type SysOperLogDeleteRes struct {
	commonApi.EmptyRes
}

type SysOperLogClearReq struct {
	g.Meta `path:"/operLog/clear" tags:"操作日志" method:"delete" summary:"清除日志"`
	commonApi.Author
}

type SysOperLogClearRes struct {
	commonApi.EmptyRes
}

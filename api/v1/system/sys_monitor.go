package system

import (
	"github.com/gogf/gf/v2/frame/g"
)

type MonitorReq struct {
	g.Meta `path:"/monitor/server" tags:"服务监控" method:"get" summary:"服务监控"`
}

type MonitorRes g.Map

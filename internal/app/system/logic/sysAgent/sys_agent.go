package sysAgent

import (
	"context"
	"ugodubai-server/api/v1/system"
	"ugodubai-server/internal/app/system/consts"
	"ugodubai-server/internal/app/system/dao"
	"ugodubai-server/internal/app/system/service"
	"ugodubai-server/library/liberr"

	"github.com/gogf/gf/frame/g"
)

func init() {
	service.RegisterSysAgent(New())
}

func New() *sSysAgent {
	return &sSysAgent{}
}

type sSysAgent struct {
}

// List 代理商列表

func (s *sSysAgent) List(ctx context.Context, req *system.AgentSearchReq) (res *system.AgentSearchRes, err error) {
	res = new(system.AgentSearchRes)
	err = g.Try(func() {
		m := dao.SysAgent.Ctx(ctx)

		res.Total, err = m.Count()
		liberr.ErrIsNil(ctx, err, "代理商列表获取失败")

		if req.PageNum == 0 {
			req.PageNum = 1
		}
		if req.PageSize == 0 {
			req.PageSize = consts.PageSize
		}

		res.CurrentPage = req.PageNum
		err = m.Page(req.PageNum, req.PageSize).Order("id asc").Scan(&res.Agents)
		liberr.ErrIsNil(ctx, err, "代理商列表获取失败")
	})
	return
}

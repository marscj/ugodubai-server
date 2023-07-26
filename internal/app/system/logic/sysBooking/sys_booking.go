package sysBooking

import (
	"context"
	"ugodubai-server/api/v1/system"
	"ugodubai-server/internal/app/system/consts"
	"ugodubai-server/internal/app/system/dao"
	"ugodubai-server/internal/app/system/model"
	"ugodubai-server/internal/app/system/service"
	"ugodubai-server/library/liberr"

	"github.com/gogf/gf/v2/frame/g"
)

func init() {
	service.RegisterSysBooking(New())
}

func New() *sSysBooking {
	return &sSysBooking{}
}

type sSysBooking struct {
}

// List 代理商列表
func (s *sSysBooking) List(ctx context.Context, req *system.BookingListReq) (total interface{}, bookingList []*model.SysBooking, err error) {

	err = g.Try(ctx, func(ctx context.Context) {
		m := dao.SysBooking.Ctx(ctx)
		total, err = m.Count()
		liberr.ErrIsNil(ctx, err, "代理商列表获取失败")

		if req.PageNum == 0 {
			req.PageNum = 1
		}
		if req.PageSize == 0 {
			req.PageSize = consts.PageSize
		}

		err = m.Page(req.PageNum, req.PageSize).Order("id desc").Scan(&bookingList)
		liberr.ErrIsNil(ctx, err, "代理商列表获取失败")
	})
	return
}

//  通过Id获取代理商信息
func (s *sSysBooking) Get(ctx context.Context, id uint64) (booking *model.SysBooking, err error) {

	err = g.Try(ctx, func(ctx context.Context) {
		err = dao.SysBooking.Ctx(ctx).WherePri(id).Scan(&booking)

		if err != nil {
			liberr.ErrIsNil(ctx, err, "代理商信息获取失败")
		}
	})
	return
}

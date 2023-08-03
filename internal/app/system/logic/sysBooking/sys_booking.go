package sysBooking

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"time"
	"ugodubai-server/api/v1/system"
	"ugodubai-server/internal/app/system/consts"
	"ugodubai-server/internal/app/system/dao"
	"ugodubai-server/internal/app/system/model"
	"ugodubai-server/internal/app/system/service"
	"ugodubai-server/library/liberr"

	"github.com/gogf/gf/os/gctx"
	"github.com/gogf/gf/v2/database/gdb"
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
		liberr.ErrIsNil(ctx, err, "代理商信息获取失败")
	})
	return
}

func (c *sSysBooking) Checkout(ctx context.Context, req *system.CheckOutReq) (res *system.CheckOutRes, err error) {

	err = g.Validator().Assoc(req).Data(req.Item).Run(gctx.New())

	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		err = g.Try(ctx, func(ctx context.Context) {

		})
		return err
	})

	secretKey := g.Cfg().MustGet(ctx, "system.secretKey")
	fmt.Println("Generated key:", secretKey)
	item := &system.PreCheckOutItem{
		Quantity:         1,
		VariationPriceId: 1,
		ActionDate:       "2020-10-10 10:10",
		Price:            "11.20",
		Tax:              "0",
		Timestamp:        time.Now(),
	}

	GenerateSignature(secretKey.String(), item)

	b := VerifySignature(secretKey.String(), item)

	fmt.Println("result:", b)

	return
}

// 加密
func GenerateSignature(secretKey string, data *system.PreCheckOutItem) error {
	dataBytes, err := json.Marshal(data)
	if err != nil {
		return err
	}

	h := hmac.New(sha256.New, []byte(secretKey))
	_, err = h.Write(dataBytes)
	if err != nil {
		return err
	}

	data.Signature = hex.EncodeToString(h.Sum(nil))
	return nil
}

// 解密
func VerifySignature(secretKey string, data *system.PreCheckOutItem) bool {
	signature := data.Signature // Save original signature
	data.Signature = ""         // Clear the signature for re-generation

	// Re-generate the signature
	err := GenerateSignature(secretKey, data)
	if err != nil {
		return false
	}

	// Compare the newly generated signature with the original one
	isEqual := (data.Signature == signature)
	data.Signature = signature // Restore the original signature

	if !isEqual {
		return false
	}

	// Check if the timestamp is within the last 30 minutes
	now := time.Now()
	if now.Sub(data.Timestamp).Minutes() > 30 {
		return false
	}

	return true
}

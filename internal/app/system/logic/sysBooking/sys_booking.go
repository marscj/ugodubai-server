package sysBooking

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
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

// 通过Id获取代理商信息
func (s *sSysBooking) Get(ctx context.Context, id uint64) (booking *model.SysBooking, err error) {

	err = g.Try(ctx, func(ctx context.Context) {
		err = dao.SysBooking.Ctx(ctx).WherePri(id).Scan(&booking)
		liberr.ErrIsNil(ctx, err, "代理商信息获取失败")
	})
	return
}

// 检查
func (c *sSysBooking) PreCheckout(ctx context.Context, req *system.PreCheckOutReq) (subTotal string, subTax string, total string, items []*model.SysPreCheckOutItemRes, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		var (
			variationPrice []*model.SysVariationPriceByCheckout
		)
		g.Model("sys_variation_price").Where("variation_price_id IN(1,2)").WithAll().Scan(&variationPrice)

		subTotal = "1000"
		subTax = "0.00"
		total = "1000.00"
	})
	return
}

// 结算
func (c *sSysBooking) Checkout(ctx context.Context, req *system.CheckOutReq) (err error) {

	err = g.Validator().Assoc(req).Data(req.Item).Run(gctx.New())

	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		err = g.Try(ctx, func(ctx context.Context) {

		})
		return err
	})

	return
}

// func (c *sSysBooking) PreCheckout(ctx context.Context, req *system.PreCheckOutReq) (subTotal string, subTax string, Total string, items []*model.SysPreCheckOutItemRes, err error) {

// 	secretKey := g.Cfg().MustGet(ctx, "system.secretKey")
// 	fmt.Println("Generated key:", secretKey)
// 	item := &model.SysPreCheckOutItemRes{
// 		Quantity:         1,
// 		VariationPriceId: 1,
// 		ActionDate:       "2020-10-10 10:10",
// 		Timestamp:        time.Now().UTC(),
// 	}

// 	GenerateSignature(secretKey.String(), item)

// 	b := VerifySignature(secretKey.String(), item)

// 	fmt.Println("result:", b)
// 	items = append(items, item)
// 	return
// }

// 加密
func GenerateSignature(secretKey string, data *model.SysPreCheckOutItemRes) error {
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

// 解密 验证签名以及时间戳 < 30 分钟
func VerifySignature(secretKey string, data *model.SysPreCheckOutItemRes) bool {
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

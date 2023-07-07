package controller

import (
	"context"
	"ugodubai-server/api/v1/common"
	"ugodubai-server/internal/app/common/service"
)

var Captcha = captchaController{}

type captchaController struct {
}

// Get 获取验证码
func (c *captchaController) Get(ctx context.Context, req *common.CaptchaReq) (res *common.CaptchaRes, err error) {
	var (
		idKeyC, base64stringC string
	)
	idKeyC, base64stringC, err = service.Captcha().GetVerifyImgString(ctx)
	res = &common.CaptchaRes{
		Key: idKeyC,
		Img: base64stringC,
	}
	return
}

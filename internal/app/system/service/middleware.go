package service

import (
	"github.com/gogf/gf/v2/net/ghttp"
)

type IMiddleware interface {
	Ctx(r *ghttp.Request)
	Auth(r *ghttp.Request)
}

var localMiddleware IMiddleware

func Middleware() IMiddleware {
	if localMiddleware == nil {
		panic("implement not found for interface IMiddleware, forgot register?")
	}
	return localMiddleware
}

func RegisterMiddleware(i IMiddleware) {
	localMiddleware = i
}

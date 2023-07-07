package controller

import (
	"context"
	"ugodubai-server/api/v1/system"
	commonConsts "ugodubai-server/internal/app/common/consts"
	"ugodubai-server/internal/app/common/service"
	"ugodubai-server/internal/app/system/consts"

	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

var Cache = new(cacheController)

type cacheController struct {
	BaseController
}

func (c *cacheController) Remove(ctx context.Context, req *system.CacheRemoveReq) (res *system.CacheRemoveRes, err error) {
	service.Cache().RemoveByTag(ctx, commonConsts.CacheSysDictTag)
	service.Cache().RemoveByTag(ctx, commonConsts.CacheSysConfigTag)
	service.Cache().RemoveByTag(ctx, consts.CacheSysAuthTag)
	cacheRedis := g.Cfg().MustGet(ctx, "system.cache.model").String()
	if cacheRedis == commonConsts.CacheModelRedis {
		cursor := 0
		cachePrefix := g.Cfg().MustGet(ctx, "system.cache.prefix").String()
		for {
			var v *gvar.Var
			v, err = g.Redis().Do(ctx, "scan", cursor, "match", cachePrefix+"*", "count", "100")
			if err != nil {
				return
			}
			data := gconv.SliceAny(v)
			var dataSlice []string
			err = gconv.Structs(data[1], &dataSlice)
			if err != nil {
				return
			}
			for _, d := range dataSlice {
				_, err = g.Redis().Do(ctx, "del", d)
				if err != nil {
					return
				}
			}
			cursor = gconv.Int(data[0])
			if cursor == 0 {
				break
			}
		}
	}
	return
}

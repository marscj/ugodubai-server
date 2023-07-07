package system

import (
	commonApi "ugodubai-server/api/v1/common"

	"github.com/gogf/gf/v2/frame/g"
)

type CacheRemoveReq struct {
	g.Meta `path:"/cache/remove" tags:"缓存管理" method:"delete" summary:"清除缓存"`
	commonApi.Author
}

type CacheRemoveRes struct {
	commonApi.EmptyRes
}

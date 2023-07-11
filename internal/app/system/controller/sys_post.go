package controller

import (
	"context"
	"ugodubai-server/api/v1/system"
	"ugodubai-server/internal/app/system/service"
)

var Post = postController{}

type postController struct {
	BaseController
}

// List 岗位列表
func (c *postController) List(ctx context.Context, req *system.PostListReq) (res *system.PostListRes, err error) {
	res, err = service.SysPost().List(ctx, req)
	return
}

// Add 添加岗位
func (c *postController) Add(ctx context.Context, req *system.PostAddReq) (res *system.PostAddRes, err error) {
	err = service.SysPost().Add(ctx, req)
	return
}

// Edit 修改岗位
func (c *postController) Edit(ctx context.Context, req *system.PostEditReq) (res *system.PostEditRes, err error) {
	err = service.SysPost().Edit(ctx, req)
	return
}

// Delete 删除岗位
func (c *postController) Delete(ctx context.Context, req *system.PostDeleteReq) (res *system.PostDeleteRes, err error) {
	err = service.SysPost().Delete(ctx, req.Ids)
	return
}

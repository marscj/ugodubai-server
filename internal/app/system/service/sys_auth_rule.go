package service

import (
	"context"

	"ugodubai-server/api/v1/system"
	"ugodubai-server/internal/app/system/model"
	"ugodubai-server/internal/app/system/model/entity"
)

type (
	ISysAuthRule interface {
		GetMenuListSearch(ctx context.Context, req *system.RuleListReq) (res []*model.SysAuthRuleInfo, err error)
		GetIsMenuList(ctx context.Context) ([]*model.SysAuthRuleInfo, error)
		GetMenuList(ctx context.Context) (list []*model.SysAuthRuleInfo, err error)
		GetIsButtonList(ctx context.Context) ([]*model.SysAuthRuleInfo, error)
		Add(ctx context.Context, req *system.RuleAddReq) (err error)
		BindRoleRule(ctx context.Context, ruleId interface{}, roleIds []uint) (err error)
		Get(ctx context.Context, id uint) (rule *entity.SysAuthRule, err error)
		GetMenuRoles(ctx context.Context, id uint) (roleIds []uint, err error)
		Update(ctx context.Context, req *system.RuleUpdateReq) (err error)
		UpdateRoleRule(ctx context.Context, ruleId uint, roleIds []uint) (err error)
		GetMenuListTree(pid uint, list []*model.SysAuthRuleInfo) []*model.SysAuthRuleTree
		DeleteMenuByIds(ctx context.Context, ids []int) (err error)
		FindSonByParentId(list []*model.SysAuthRuleInfo, pid uint) []*model.SysAuthRuleInfo
	}
)

var (
	localSysAuthRule ISysAuthRule
)

func SysAuthRule() ISysAuthRule {
	if localSysAuthRule == nil {
		panic("implement not found for interface ISysAuthRule, forgot register?")
	}
	return localSysAuthRule
}

func RegisterSysAuthRule(i ISysAuthRule) {
	localSysAuthRule = i
}

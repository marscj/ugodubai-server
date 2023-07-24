// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// SysPaymentTokens is the golang structure for table sys_payment_tokens.
type SysPaymentTokens struct {
	TokenId   int64  `json:"tokenId"   description:""`
	GatewayId string `json:"gatewayId" description:""`
	Token     string `json:"token"     description:""`
	UserId    uint64 `json:"userId"    description:""`
	AgentId   uint64 `json:"agentId"   description:""`
	Type      string `json:"type"      description:""`
	IsDefault int    `json:"isDefault" description:""`
}

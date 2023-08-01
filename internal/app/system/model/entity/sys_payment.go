// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// SysPayment is the golang structure for table sys_payment.
type SysPayment struct {
	PaymentId uint64 `json:"paymentId" description:""`
	Token     string `json:"token"     description:""`
	GatewayId string `json:"gatewayId" description:""`
	BookingId uint64 `json:"bookingId" description:""`
}

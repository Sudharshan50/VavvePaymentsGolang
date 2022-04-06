package src

import (
	"time"
)

type Payment struct {
	MerchantId          string    `json:"merchantId"`
	PaymentThru         string    `json:"paymentThru"`
	PaymentInfo         string    `json:"paymentInfo"`
	PaymentAuthId       string    `json:"paymentAuthId"`
	PaymentStatus       string    `json:"paymentStatus"`
	StoreDiscType       string    `json:"storeDiscType"`
	Amount              float32   `json:"amount"`
	TotalAmount         float32   `json:"totalAmount"`
	PaymentReferrenceId string    `json:"paymentReferrenceId"`
	CreatedBy           string    `json:"createdBy"`
	CreatedDate         time.Time `json:"createdDate"`
	UpdatedBy           string    `json:"updatedBy"`
	UpdatedDate         time.Time `json:"updatedDate"`
	Description         string    `json:"description"`
	MerchantLocationId  int64     `json:"merchantLocationId"`
	MemberId            string    `json:"memberId"`
	RedeemType          string    `json:"redeemType"`
	OrderId             string    `json:"orderId"`
	OrderType           string    `json:"orderType"`
}

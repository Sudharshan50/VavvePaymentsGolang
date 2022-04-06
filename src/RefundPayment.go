package src

type RefundPayment struct {
	PaymentTransactionId int64   `json:"paymentTransactionId"`
	MerchantId           string  `json:"merchantId"`
	RefundAmount         float32 `json:"refundAmount"`
	Currency             string  `json:"currency"`
	PaymentMethodType    string  `json:"paymentMethodType"`
}

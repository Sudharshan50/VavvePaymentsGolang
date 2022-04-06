package src

type CapturePayment struct {
	PaymentTransactionId int64  `json:"paymentTransactionId"`
	MerchantId           string `json:"merchantId"`
	PaymentMethodType    string `json:"paymentMethodType"`
}

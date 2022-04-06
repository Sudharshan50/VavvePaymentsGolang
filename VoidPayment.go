package VavvePaymentsGolang

type VoidPayment struct {
	PaymentTransactionId int64   `json:"paymentTransactionId"`
	MerchantId           string  `json:"merchantId"`
	VoidPayment          float32 `json:"voidPayment"`
	Currency             string  `json:"currency"`
	PaymentMethodType    string  `json:"PaymentMethodType"`
}

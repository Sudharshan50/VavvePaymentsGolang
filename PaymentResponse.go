package VavvePaymentsGolang

type PayemntResponse struct {
	PaymentTransactionId       int64   `json:"paymentTransactionId"`
	MerchantOrderId            int64   `json:"merchantOrderId"`
	TotalOrderAmount           float32 `json:"totalOrderAmount"`
	RedirectURL                string  `json:"redirectURL"`
	Amount                     float32 `json:"amount"`
	Currency                   string  `json:"currency"`
	TransactionId              string  `json:"transactionId"`
	Status                     string  `json:"status"`
	HttpResponseCode           int16   `json:"httpResponseCode"`
	HttpResponseReason         string  `json:"httpResponseReason"`
	ErrorReason                string  `json:"errorReason"`
	ErrorMessage               string  `json:"errorMessage"`
	ErrorField                 string  `json:"errorField"`
	ErrorFieldReason           string  `json:"errorFieldReason"`
	InformationApprovalCode    int64   `json:"informationApprovalCode"`
	InformationResponseCode    int16   `json:"informationResponseCode"`
	AvsCode                    string  `json:"avsCode"`
	AvsCodeRaw                 string  `json:"avsCodeRaw"`
	ReconciliationId           string  `json:"reconciliationId"`
	OrderSubmitTimeUtc         string  `json:"orderSubmitTimeUtc"`
	CardVerificationResultCode string  `json:"cardVerificationResultCode"`
	CvvStatus                  string  `json:"cvvStatus"`
	RefundId                   string  `json:"refundId"`
	ProductId                  int64   `json:"productId"`
}

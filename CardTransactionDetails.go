package VavvePaymentsGolang

type CardTransactionDetails struct {
	PaymentTransactionId           int64   `json:"paymentTransactionId"`
	MerchantId                     string  `json:"merchantId"`
	MerchantOrderId                int64   `json:"merchantOrderId"`
	IpAddress                      string  `json:"ipAddress"`
	MachineName                    string  `json:"machineName"`
	RequestURL                     string  `json:"requestURL"`
	CreatedDate                    string  `json:"createdDate"`
	TotalOrderAmount               float32 `json:"totalOrderAmount"`
	CardNumberLast4Digit           string  `json:"cardNumberLast4Digit"`
	Currency                       string  `json:"currency"`
	Amount                         float32 `json:"amount"`
	PaymentType                    string  `json:"paymentType"`
	AvsCode                        string  `json:"avsCode"`
	AvsCodeRaw                     string  `json:"avsCodeRaw"`
	CardSubTypeId                  int64   `json:"cardSubTypeId"`
	CardTypeId                     string  `json:"cardTypeId"`
	CardVerificationResultCode     string  `json:"cardVerificationResultCode"`
	CvvStatus                      string  `json:"cvvStatus"`
	ErrorReason                    string  `json:"errorReason"`
	ErrorMessage                   string  `json:"errorMessage"`
	ErrorField                     string  `json:"errorField"`
	ErrorFieldReason               string  `json:"errorFieldReason"`
	HttpResponseCode               int16   `json:"httpResponseCode"`
	HttpResponseReason             string  `json:"httpResponseReason"`
	ProcessingStartTime            string  `json:"processingStartTime"`
	ProcessingEndTime              string  `json:"processingEndTime"`
	Status                         string  `json:"status"`
	InformationApprovalCode        int64   `json:"informationApprovalCode"`
	InformationResponseCode        int16   `json:"informationResponseCode"`
	OrderSubmitTimeUtc             string  `json:"orderSubmitTimeUtc"`
	ReconciliationId               string  `json:"reconciliationId"`
	ProcessorRefundId              string  `json:"processorRefundId"`
	TransactionId                  string  `json:"transactionId"`
	CapturePaymentResponseAmount   float32 `json:"capturePaymentResponseAmount"`
	CapturePaymentResponseCurrency string  `json:"capturePaymentResponseCurrency"`
	OldCardPaymentTransactionId    int64   `json:"oldCardPaymentTransactionId"`
	GeoLocationId                  string  `json:"geoLocationId"`
	PaymentThru                    string  `json:"paymentThru"`
	PaymentApp                     string  `json:"paymentApp"`
	PaymentAppType                 string  `json:"paymentAppType"`
	ModeOfPayment                  string  `json:"modeOfPayment"`
}

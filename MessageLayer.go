package src

type MessageLayer struct {
	Msgidentifier      string  `json:"msgidentifier"`
	MerchantId         string  `json:"merchantId"`
	OrderId            string  `json:"orderId"`
	TotalOrderAmount   float32 `json:"totalOrderAmount"`
	Amount             float32 `json:"amount"`
	Currency           string  `json:"currency"`
	Account            string  `json:"account"`
	Holdername         string  `json:"holdername"`
	Bankname           string  `json:"bankname"`
	BankId             int64   `json:"bankId"`
	Accounttype        string  `json:"accounttype"`
	Routingnumber      string  `json:"routingnumber"`
	Cvv2               string  `json:"cvv2"`
	Expiry             string  `json:"expiry"`            // format YYMM
	Modeofpayment      string  `json:"modeofpayment"`     // online or offline
	PaymentApp         string  `json:"paymentApp"`        // Appointment/Giftcard/etc
	PaymentAppType     string  `json:"paymentAppType"`    // Hosted/Self/API
	PaymentMethodType  string  `json:"paymentMethodType"` // card/ach/digital
	PaymentThru        string  `json:"paymentThru"`       // online or vavve
	RequestURL         string  `json:"requestURL"`
	Track              string  `json:"track"`
	Address1           string  `json:"address1"`
	Address2           string  `json:"address2"`
	City               string  `json:"city"`
	State              string  `json:"state"`
	Country            string  `json:"country"`
	Zipcode            string  `json:"zipcode"`
	Customerid         string  `json:"customerid"`
	Firstname          string  `json:"firstname"`
	Lastname           string  `json:"lastname"`
	Phone              string  `json:"phone"`
	Email              string  `json:"email"`
	ExitingStatus      string  `json:"exitingStatus"`
	MemberId           string  `json:"memberId"`
	MerchantLocationId int64   `json:"merchantLocationId"`
}

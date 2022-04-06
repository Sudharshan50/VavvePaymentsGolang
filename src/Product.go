package src

type Product struct {
	OrderItemId     int64   `json:"orderItemId"`
	ProductId       string  `json:"productId"`
	ProductName     string  `json:"productName"`
	Price           float32 `json:"price"`
	Quantity        int32   `json:"quantity"`
	Itemdescription string  `json:"itemdescription"`
}

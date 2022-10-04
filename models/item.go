package models

type Item struct {
	Item_id     int    `json:"itemsId" gorm:"primary_key;auto_increment;not_null"`
	Item_code   string `json:"itemCode"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
	Order_id    int    `json:"orderId"`
}

type RequestItem struct {
	Item_code   string `json:"itemCode"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
}

type RequestUpdateItem struct {
	Item_id     int    `json:"lineItemId"`
	Item_code   string `json:"itemCode"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
}

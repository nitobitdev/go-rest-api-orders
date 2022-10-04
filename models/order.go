package models

import "time"

type Order struct {
	Order_id      int       `json:"orderId" gorm:"primary_key;auto_increment;not_null"`
	Customer_name string    `json:"customerName"`
	Ordered_at    time.Time `json:"orderedAt"`
}

type RequestUpdateOrder struct {
	Customer_name string    `json:"customerName"`
	Ordered_at    time.Time `json:"orderedAt"`
}

type RequestOrder struct {
	Ordered_at    time.Time     `json:"orderedAt"`
	Customer_name string        `json:"customerName"`
	Item          []RequestItem `json:"items"`
}

type ResponseOrder struct {
	Order_id      int       `json:"orderId"`
	Customer_name string    `json:"customerName"`
	Ordered_at    time.Time `json:"orderedAt"`
}

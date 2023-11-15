package data

import "time"

type Customer struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Order struct {
	OrderID     int           `json:"order_id"`
	Customer    Customer      `json:"customer"`
	OrderNumber string        `json:"order_number"`
	CreatedAt   *time.Time    `json:"created_at"`
	Tax         int           `json:"tax"`
	Discount    int           `json:"discount"`
	Total       int           `json:"total"`
	Details     []OrderDetail `json:"details"`
}

type OrderProduct struct {
	ProductID int    `json:"product_id"`
	Name      string `json:"name"`
	Price     int    `json:"price"`
}

type OrderDetail struct {
	OrderDetailID int          `json:"order_detail_id"`
	OrderID       int          `json:"order_id"`
	Product       OrderProduct `json:"product"`
	Qty           int          `json:"qty"`
	SubTotal      int          `json:"sub_total"`
}

type CreateOrderRequest struct {
	OrderNumber string               `json:"order_number"`
	Details     []OrderDetailRequest `json:"details"`
}

type OrderDetailRequest struct {
	Product OrderProduct `json:"product"`
	Qty     int          `json:"qty"`
}

var OrderListData = []*Order{
	{
		OrderID:     1,
		OrderNumber: "ORD-001",
		Customer: Customer{
			Name:  "Joshua",
			Email: "antoniusjoshua47@gmail.com",
		},
		CreatedAt: nil,
		Tax:       0,
		Discount:  0,
		Total:     9000,
		Details: []OrderDetail{
			{
				OrderDetailID: 1,
				OrderID:       1,
				Product: OrderProduct{
					ProductID: 1,
					Name:      "Shampoo",
					Price:     1000,
				},
				Qty:      1,
				SubTotal: 1000,
			},
			{
				OrderDetailID: 2,
				OrderID:       1,
				Product: OrderProduct{
					ProductID: 2,
					Name:      "Ketchup",
					Price:     2000,
				},
				Qty:      4,
				SubTotal: 8000,
			},
		},
	},
}

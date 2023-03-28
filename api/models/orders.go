package models

type Order struct {
	Id        		string  `json:"id"`
	Name      		string  `json:"name"`
	Price    		float64 	`json:"price"`
	Phone_number	string	`json:"phone_number"`
	Latitude		int		`json:"latitude"`
	Longtitude		int		`json:"longtitude"`
	User_id			string	`json:"user_id"`
	Customer_id		string	`json:"customer_id"`
	Courier_id		string	`json:"courier_id"`
	Product_id		string	`json:"product_id"`
	Quantity		int		`json:"quantity"`
	CreatedAt 		string  `json:"created_at"`
	UpdatedAt 		string  `json:"updated_at"`
}

type OrderPrimaryKey struct {
	Id string `json:"id"`
}

type CreateOrderSwagger struct {
	Name      		string  `json:"name"`
	Phone_number	string	`json:"phone_number"`
	Latitude		int		`json:"latitude"`
	Longtitude		int		`json:"longtitude"`
	User_id			string	`json:"user_id"`
	Customer_id		string	`json:"customer_id"`
	Courier_id		string	`json:"courier_id"`
	Product_id		string	`json:"product_id"`
	Quantity		int		`json:"quantity"`
}

type CreateOrder struct {
	Name      		string  `json:"name"`
	Price    		float64 	`json:"price"`
	Phone_number	string	`json:"phone_number"`
	Latitude		int		`json:"latitude"`
	Longtitude		int		`json:"longtitude"`
	User_id			string	`json:"user_id"`
	Customer_id		string	`json:"customer_id"`
	Courier_id		string	`json:"courier_id"`
	Product_id		string	`json:"product_id"`
	Quantity		int		`json:"quantity"`
}

type UpdateOrder struct {
	Id        		string  `json:"id"`
	Name      		string  `json:"name"`
	Price    		float64 	`json:"price"`
	Phone_number	string	`json:"phone_number"`
	Latitude		int		`json:"latitude"`
	Longtitude		int		`json:"longtitude"`
	User_id			string	`json:"user_id"`
	Customer_id		string	`json:"customer_id"`
	Courier_id		string	`json:"courier_id"`
	Product_id		string	`json:"product_id"`
	Quantity		int		`json:"quantity"`
}

type GetListOrderRequest struct {
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
	Search string `json:"search"`
}

type GetListOrderResponse struct {
	Count 	int     	`json:"count"`
	Orders	[]*Order	`json:"orders"`
}
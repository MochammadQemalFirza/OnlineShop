package web

import "gopkg.in/guregu/null.v4"


type Products struct {
	ProductID    null.Int      `json:"product_id"`
	CategoryName null.String   `json:"category_name"`
	StoreName    null.String   `json:"store_name"`
	ProductName  null.String   `json:"product_name"`
	Quantity     null.Int      `json:"quantity"`
	Price        null.String   `json:"price"`
}

package domain

import "gopkg.in/guregu/null.v4"
type User struct {
	UserID   null.String `json:"user_id"`
	Username null.String `json:"username"`
	Password null.String `json:"passwords"`
	Address  null.String `json:"address"`
	Age      null.Int    `json:"age"`
}

type Product struct {
	ProductID   null.Int     `json:"product_id"`
	Category    Category     `json:"category"`
	Store       Store        `json:"store"`
	ProductName null.String  `json:"product_name"`
	Quantity    null.Int     `json:"quantity"`
	Price       null.Int      `json:"price"`
}

type Category struct {
	CategoryID   null.String `json:"category_id"`
	CategoryName null.String `json:"category_name"`
}

type Store struct {
	StoreID   null.String `json:"store_id"`
	StoreName null.String `json:"store_name"`
}

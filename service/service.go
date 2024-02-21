package service

import (
	"context"

	"github.com/MochammadQemalFirza/OnlineShop/model/web"
)

type Service interface {
	GetListProduct(ctx context.Context, category string) ([]web.Products, error)
}
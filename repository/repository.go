package repository

import (
	"context"

	"github.com/MochammadQemalFirza/OnlineShop/model/domain"
)

type Repository interface {
	GetListProduct(ctx context.Context) (ListProduct []domain.Product, err error)
}
package service

import (
	"context"
	"strconv"

	"github.com/MochammadQemalFirza/OnlineShop/model/domain"
	"github.com/MochammadQemalFirza/OnlineShop/model/web"
	repository "github.com/MochammadQemalFirza/OnlineShop/repository"
	"gopkg.in/guregu/null.v4"
)

type ServiceImpl struct {
	Repository repository.Repository
}

func (s *ServiceImpl)GetListProduct(ctx context.Context, category string) ([]web.Products, error){
	listProduct, err := s.Repository.GetListProduct(ctx)
	if err != nil {
		return nil, err
	}
	filteredProducts := FilterProductsByCategory(listProduct, category)
    listProductResponse := ToGetListProductResponse(filteredProducts)
 
	return listProductResponse, nil
}

func convInt (s int) string{
	conv := s
	convert := strconv.Itoa(conv)
	return convert
} 

func ToGetListProductResponse(listProduct []domain.Product) (listProductResponse []web.Products) {
	for _, product := range listProduct {
		listProductResponse = append(listProductResponse, web.Products{
			ProductID     :   product.ProductID,
			CategoryName  :   product.Category.CategoryName,
			StoreName     :   product.Store.StoreName,
			ProductName   :   product.ProductName,
			Price         :   null.StringFrom("Rp. " + convInt(int(product.Price.Int64))),
			Quantity      :   product.Quantity,
		},
		)
	}

	return listProductResponse
}

func FilterProductsByCategory(products []domain.Product, category string) []domain.Product {
	var filteredProducts []domain.Product
	for _, p := range products {
		if category != "" && p.Category.CategoryName.String != category {
			continue
		}
		filteredProducts = append(filteredProducts, p)
	}
	return filteredProducts
}

func NewService(Repository repository.Repository) Service {
	return &ServiceImpl{Repository: Repository}
}
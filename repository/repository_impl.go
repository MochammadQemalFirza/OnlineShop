package repository

import (
	"context"
	"database/sql"

	serverError "github.com/MochammadQemalFirza/OnlineShop/helpers/errors/server"
	"github.com/MochammadQemalFirza/OnlineShop/model/domain"
)
type RepositoryImpl struct {
	db *sql.DB
}

func (r *RepositoryImpl) GetListProduct(ctx context.Context) (ListProduct []domain.Product, err error) {
	var (
		values       []interface{}
		query        string
	)

	query =`SELECT 
    p.product_id,
    c.category_name,
    s.store_name,
    p.product_name,
    p.price,
    p.quantity
FROM 
    products p
LEFT JOIN 
    category c ON p.categories_id = c.category_id
LEFT JOIN 
    store s ON p.store_id = s.store_id;`

	stmt, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return nil, serverError.NewDataSourceError(err, "error preparing query")
	}
	defer stmt.Close()

	result, err := stmt.QueryContext(ctx, values...)
	if err != nil {
		return nil, serverError.NewDataSourceError(err, "error executing query")
	}
	defer result.Close()

	for result.Next() {
		product := domain.Product{}
		err := result.Scan(
			&product.ProductID,
			&product.Category.CategoryName,
			&product.Store.StoreName,
			&product.ProductName,
			&product.Price,
			&product.Quantity,
		)
		if err != nil {
			return nil, serverError.NewDataSourceError(err, "error scanning row")
		}

		ListProduct = append(ListProduct, product)
	}
	return ListProduct, nil
}


func NewRepository(db *sql.DB) Repository {
	return &RepositoryImpl{db: db}
}
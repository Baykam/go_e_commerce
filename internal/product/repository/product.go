package repository

import (
	"context"
	"golang_testing_grpc/internal/product/dto"
	"golang_testing_grpc/internal/product/model"
	"golang_testing_grpc/pkg/config"
	"golang_testing_grpc/pkg/db"
	"golang_testing_grpc/pkg/paging"

	"github.com/quangdangfit/gocommon/logger"
)

type IProductRepository interface {
	ListProducts(ctx context.Context, req *dto.ListProductReq, page int, limit int) ([]*model.Product, *paging.Pagination, error)
}

type ProductRepo struct {
	db db.IDatabaseInterface
}

func NewProductRepository(db db.IDatabaseInterface) *ProductRepo {
	return &ProductRepo{
		db: db,
	}
}

func (r *ProductRepo) ListProducts(ctx context.Context, req *dto.ListProductReq, page int, limit int) ([]*model.Product, *paging.Pagination, error) {
	ctx, cancel := context.WithTimeout(ctx, config.DatabaseTimeOut)
	defer cancel()
	rows, err := r.db.JustQuery(ctx, db.ProductTable, r.db.GetLimit(limit))
	if err != nil {
		return nil, nil, err
	}
	defer rows.Close()

	var productList []*model.Product
	for rows.Next() {
		var product model.Product
		if err := rows.Scan(
			&product.ID,
			&product.CreatedAt,
			&product.UpdatedAt,
			&product.DeletedAt,
			&product.Code,
			&product.Name,
			&product.Description,
			&product.Price,
			&product.Active,
		); err != nil {
			return nil, nil, err
		}
		productList = append(productList, &product)
	}

	if err := rows.Err(); err != nil {
		logger.Error("Error during rows iteration:", err)
		return nil, nil, err
	}
	totalItemCount := 100
	pagination := paging.New(1, 10, int64(totalItemCount))
	return productList, pagination, nil
}
func (r *ProductRepo) GetBestDeals() {

}

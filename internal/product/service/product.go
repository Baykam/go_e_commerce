package service

import (
	"context"
	"golang_testing_grpc/internal/product/dto"
	"golang_testing_grpc/internal/product/model"
	"golang_testing_grpc/internal/product/repository"
	"golang_testing_grpc/pkg/paging"

	"github.com/quangdangfit/gocommon/validation"
)

type IProductService interface {
	ListProducts(ctx context.Context, req *dto.ListProductReq, limit int, page int) ([]*model.Product, *paging.Pagination, error)
}

type ProductService struct {
	validator validation.Validation
	repo      repository.IProductRepository
}

func NewProductService(validator validation.Validation, repo repository.IProductRepository) *ProductService {
	return &ProductService{
		validator: validator,
		repo:      repo,
	}
}

func (s *ProductService) ListProducts(ctx context.Context, req *dto.ListProductReq, limit int, page int) ([]*model.Product, *paging.Pagination, error) {
	productList, pagination, err := s.repo.ListProducts(ctx, req, page, limit)
	if err != nil {
		return nil, nil, err
	}

	return productList, pagination, nil
}

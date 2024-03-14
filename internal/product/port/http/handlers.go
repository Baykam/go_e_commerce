package http

import (
	"golang_testing_grpc/internal/product/dto"
	"golang_testing_grpc/internal/product/service"
	"golang_testing_grpc/pkg/paging"
	"golang_testing_grpc/pkg/response"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/quangdangfit/gocommon/logger"
)

type ProductHandlers struct {
	service      service.IProductService
	lastPosition struct {
		sync.Mutex
		Page  int64
		Limit int64
	}
}

func NewProductHandlers(service service.IProductService) *ProductHandlers {
	return &ProductHandlers{
		service: service,
	}
}

func (h *ProductHandlers) ListProducts(c *gin.Context) {
	var req dto.ListProductReq

	page := paging.ApiConverterQuery("page", c, h.lastPosition.Page)
	limit := paging.ApiConverterQuery("limit", c, h.lastPosition.Limit)

	var res dto.ListProductRes
	productList, pagination, err := h.service.ListProducts(c, &req, int(limit), int(page))
	if err != nil {
		logger.Error("Failed to list products: ", err)
		response.Error(c, http.StatusInternalServerError, err, "Failed to list products")
		return
	}

	for _, p := range productList {
		res.Products = append(res.Products, &dto.Product{
			ID:          p.ID,
			Code:        p.Code,
			Name:        p.Name,
			Description: p.Description,
			Price:       p.Price,
			Active:      p.Active,
			CreatedAt:   p.CreatedAt,
			UpdatedAt:   p.UpdatedAt,
		})
	}

	res.Pagination = *pagination

	response.JSON(c, http.StatusOK, res)
}

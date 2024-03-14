package http

import (
	"golang_testing_grpc/internal/product/repository"
	"golang_testing_grpc/internal/product/service"
	"golang_testing_grpc/pkg/db"

	"github.com/gin-gonic/gin"
	"github.com/quangdangfit/gocommon/validation"
)

func Routes(r *gin.RouterGroup, db db.IDatabaseInterface, validator validation.Validation) {
	productRepo := repository.NewProductRepository(db)
	productSvc := service.NewProductService(validator, productRepo)
	productHandler := NewProductHandlers(productSvc)

	// authMiddleWare := middleware.JWTAuth()

	productRoute := r.Group("/products")
	{
		productRoute.GET("", productHandler.ListProducts)
	}
}

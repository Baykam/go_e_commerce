package http

import (
	"fmt"
	productHttp "golang_testing_grpc/internal/product/port/http"
	"golang_testing_grpc/pkg/config"
	"golang_testing_grpc/pkg/db"
	"golang_testing_grpc/pkg/response"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/quangdangfit/gocommon/logger"
	"github.com/quangdangfit/gocommon/validation"
)

type Server struct {
	engine    *gin.Engine
	cfg       *config.Schema
	validator validation.Validation
	db        db.IDatabaseInterface
}

func NewServer(validator validation.Validation, db db.IDatabaseInterface) *Server {
	return &Server{
		engine:    gin.Default(),
		cfg:       config.GetConfig(),
		validator: validator,
		db:        db,
	}
}

func (s Server) Run() error {
	_ = s.engine.SetTrustedProxies(nil)

	if err := s.MapRoutes(); err != nil {
		log.Fatalf("MapRoutes Error: %v", err)
	}

	// s.engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	s.engine.GET("/health", func(ctx *gin.Context) {
		response.JSON(ctx, http.StatusOK, nil)
	})

	logger.Info("Http server is listening on PORT: ", s.cfg.HttpPort)
	if err := s.engine.Run(fmt.Sprintf(":%d", s.cfg.HttpPort)); err != nil {
		log.Fatalf("Running HTTP server: %v", err)
	}

	return nil
}

func (s Server) GetEngine() *gin.Engine {
	return s.engine
}

func (s Server) MapRoutes() error {
	v1 := s.engine.Group("/api/v1")
	productHttp.Routes(v1, s.db, s.validator)
	return nil
}

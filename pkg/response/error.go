package response

import (
	"golang_testing_grpc/pkg/config"

	"github.com/gin-gonic/gin"
)

func Error(c *gin.Context, status int, err error, message string) {
	cfg := config.GetConfig()
	errorRes := map[string]any{
		"message": message,
	}

	if cfg.Environment != config.ProductEnv {
		errorRes["debug"] = err.Error()
	}

	c.JSON(status, Response{
		Error: errorRes,
	})
}

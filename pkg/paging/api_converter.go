package paging

import (
	"golang_testing_grpc/pkg/response"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/quangdangfit/gocommon/logger"
)

func ApiConverterQuery(api string, c *gin.Context, newInt int64) int64 {
	your, err := strconv.ParseInt(c.DefaultQuery(api, strconv.FormatInt(newInt, 10)), 10, 64)
	if err != nil {
		logger.Error("Failed to parse 'limit' parameter: ", err)
		response.Error(c, http.StatusBadRequest, err, "Invalid 'limit' parameter")
		return 0
	}
	return your
}

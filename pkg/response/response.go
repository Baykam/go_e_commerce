package response

import "github.com/gin-gonic/gin"

type Response struct {
	Result any `json:"result"`
	Error  any `json:"error"`
}

func JSON(c *gin.Context, status int, data any) {
	c.JSON(status, Response{
		Result: data,
	})
}

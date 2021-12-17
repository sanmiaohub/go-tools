package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/sanmiaohub/go-tools/ctx"
	"net/http"
)

type ApiResponse struct {
	RID     string      `json:"rid"`
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func Success(c *gin.Context, v interface{}) {
	cc := ctx.Wrap(c)
	resp := ApiResponse{
		RID:     cc.ID(),
		Code:    200,
		Message: "success",
		Data:    v,
	}
	c.JSON(http.StatusOK, resp)
}

func Error(c *gin.Context, err error, status ...int) {
	cc := ctx.Wrap(c)
	resp := ApiResponse{
		RID:     cc.ID(),
		Code:    200,
		Message: err.Error(),
		Data:    map[string]interface{}{},
	}
	if e, ok := err.(*XError); ok {
		resp.Code = e.Code()
	}

	httpCode := http.StatusOK
	if len(status) > 0 {
		httpCode = status[0]
	}
	c.JSON(httpCode, resp)
}

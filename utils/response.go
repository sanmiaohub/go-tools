package utils

import (
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"net/http"
)

type ApiResponse struct {
	RID     string      `json:"rid"`
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func Success(c *gin.Context, v interface{}) {
	resp := ApiResponse{
		RID:     rid(c),
		Code:    200,
		Message: "success",
		Data:    v,
	}
	c.JSON(200, resp)
}

func Error(c *gin.Context, err error, status ...int) {
	resp := ApiResponse{
		RID:     rid(c),
		Code:    200,
		Message: err.Error(),
		Data:    map[string]interface{}{},
	}
	if e, ok := err.(*XError); ok {
		resp.Code = e.Code()
	}

	statusCode := http.StatusOK
	if len(status) > 0 {
		statusCode = status[0]
	}
	c.JSON(statusCode, resp)
}

func rid(c *gin.Context) string {
	id := c.Value("x-request-id")
	if id == nil {
		id = uuid.NewV4().String()
	}
	return id.(string)
}

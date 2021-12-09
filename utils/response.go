package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ApiResponse struct {
	RID     string      `json:"rid"`
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func Success(c *gin.Context, v interface{}) {
        xid,ok := c.Get("x-request-id")
		if !ok {
			ut := time.Now().UnixMicro()
			m := md5.New()
			m.Write([]byte(strconv.Itoa(int(ut))))
			xid = m.Sum([]byte(""))
		}
	resp := ApiResponse{
		RID:     xid,
		Code:    200,
		Message: "message",
		Data:    v,
	}
	c.JSON(200, resp)
}

func Error(c *gin.Context, err error, status ...int) {
	resp := ApiResponse{
		RID:     c.Value("x-request-id").(string),
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

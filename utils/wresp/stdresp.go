package wresp

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	respCodeSuccess = 0
	respCodeFailed  = -1
	respCodeAbort   = -2
)

type response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func newResponse(code int, message string, data interface{}) *response {
	return &response{
		Code:    code,
		Message: message,
		Data:    data,
	}
}

func OK(c *gin.Context, data interface{}) {
	resp := newResponse(respCodeSuccess, "成功", data)
	c.JSON(http.StatusOK, resp)
}

func OKWithMsg(c *gin.Context, message string, data interface{}) {
	resp := newResponse(respCodeSuccess, message, data)
	c.JSON(http.StatusOK, resp)
}

func Fail(c *gin.Context, statusCode int, message string) {
	resp := newResponse(respCodeFailed, message, nil)
	c.JSON(statusCode, resp)
}

func FailWithData(c *gin.Context, statusCode int, message string, data interface{}) {
	resp := newResponse(respCodeFailed, message, data)
	c.JSON(statusCode, resp)
}

// 使用在Gin的中间件里
func Abort(c *gin.Context, statusCode int, message string, data interface{}) {
	resp := newResponse(respCodeAbort, message, data)
	c.AbortWithStatusJSON(statusCode, resp)
}

package response

import (
	"easytodo/global/consts"
	"easytodo/model/response/errcode"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code      errcode.ErrCode `json:"code"`
	Message   string          `json:"message"`
	Data      interface{}     `json:"data"`
	RequestId string          `json:"request_id"`
}

func result(c *gin.Context, ec errcode.ErrCode, data interface{}) {
	rid := c.Request.Header.Get(consts.RequestIdKey)
	c.JSON(http.StatusOK, Response{
		Code:      ec,
		Message:   ec.String(),
		Data:      data,
		RequestId: rid,
	})
}

func Success(c *gin.Context, data interface{}) {
	result(c, errcode.Success, data)
}

func Fail(c *gin.Context, ec errcode.ErrCode) {
	result(c, ec, nil)
}

package response

import (
	"easytodo/global/consts"
	"easytodo/model/response/errcode"
	"errors"
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

func Ok(c *gin.Context) {
	result(c, errcode.Success, true)
}

func OkWithData(c *gin.Context, data interface{}) {
	result(c, errcode.Success, data)
}

func Fail(c *gin.Context, err error) {
	var ec errcode.ErrCode
	// 判断是否为自定义错误
	if errors.As(err, &ec) {
		result(c, ec, nil)
	} else {
		result(c, errcode.ServerError, nil)
	}
}

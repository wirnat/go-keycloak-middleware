package response

import (
	"github.com/gin-gonic/gin"
)

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

const (
	ERROR        = 500
	SUCCESS      = 200
	UNAUTHORIZED = 401
)

func Result(code int, data interface{}, msg string, c *gin.Context) {
	// RETURN JSON RESPONSE

	c.JSON(code, Response{
		code,
		data,
		msg,
	})

}

func Ok(c *gin.Context) {
	Result(SUCCESS, map[string]interface{}{}, "success", c)
}

func OkWithMessage(message string, c *gin.Context) {
	Result(SUCCESS, map[string]interface{}{}, message, c)
}

func OkWithData(data interface{}, c *gin.Context) {
	Result(SUCCESS, data, "success", c)
}

func OkWithDetailed(data interface{}, message string, c *gin.Context) {
	Result(SUCCESS, data, message, c)
}

func Fail(c *gin.Context) {
	Result(ERROR, map[string]interface{}{}, "failed", c)
}

func UnauthorizedFailDetail(data interface{}, message string, c *gin.Context) {
	Result(UNAUTHORIZED, data, message, c)
	c.Abort()
}

func FailWithMessage(message string, c *gin.Context) {
	Result(ERROR, map[string]interface{}{}, message, c)
	c.Abort()

}

func FailWithDetailed(data interface{}, message string, c *gin.Context) {
	Result(ERROR, data, message, c)
	c.Abort()
}

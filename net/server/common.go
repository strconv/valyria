package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	SUCCESS_CODE  = 200  // 成功
	INVALID_PARAM = 499  // 参数错误
	SYS_ERROR     = 500  // 系统错误
	JWT_FAIL      = 4001 // JWT错误
)

type ApiResult struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type RespModel struct {
	ApiResult
	Data interface{} `json:"data,omitempty"`
}

func RespSuccessMsg(c *gin.Context, obj interface{}) {
	c.JSON(http.StatusOK, RespModel{
		ApiResult: ApiResult{
			Code: SUCCESS_CODE,
			Msg:  "操作成功",
		},
		Data: obj,
	})
}

func RespSuccess(c *gin.Context) {
	c.JSON(http.StatusOK, ApiResult{
		Code: SUCCESS_CODE,
		Msg:  "操作成功",
	})
}

func RespInvalidParam(c *gin.Context, msg string) {
	c.JSON(http.StatusOK, ApiResult{
		Code: INVALID_PARAM,
		Msg:  "参数错误: " + msg,
	})
}

func RespSysError(c *gin.Context) {
	c.JSON(http.StatusOK, ApiResult{
		Code: SYS_ERROR,
		Msg:  "系统错误",
	})
}

func RespBusiness(c *gin.Context, code int, msg string) {
	c.JSON(http.StatusOK, ApiResult{
		Code: code,
		Msg:  msg,
	})
}

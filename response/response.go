package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type successStruct struct {
	Code int         `json:"code"`
	Msg  interface{} `json:"msg"`
	Data interface{} `json:"data"`
}

type failureStruct struct {
	Code int         `json:"code"`
	Msg  interface{} `json:"msg"`
}

// GinSuccess 请求成功返回
func GinSuccess(c *gin.Context, data interface{}) {
	json := &successStruct{Code: successCode, Msg: successMsg, Data: data}

	c.JSON(http.StatusOK, json)
}

// GinFailure 请求失败返回
func GinFailure(c *gin.Context, code int, data interface{}) {
	json := &failureStruct{Code: code, Msg: data}

	c.JSON(http.StatusOK, json)
}

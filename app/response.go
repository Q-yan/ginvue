package app

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type Gin struct {
	C *gin.Context
}

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"message"`
	Data interface{} `json:"data"`
}

// ConcatStrings 使用Strings.builder高效拼接字符串
func ConcatStrings(str ...string) string {
	var builder strings.Builder
	for _, s := range str {
		builder.WriteString(s)
	}
	return builder.String()
}

// Response setting gin.JSON
func (g *Gin) Response(errCode int, msg string, data interface{}) {
	g.C.JSON(http.StatusOK, Response{
		Code: errCode,
		Msg:  msg,
		Data: data,
	})
	return
}

// 填写成功的response信息，data为空，code是200
func (g *Gin) ResponseSuc(msg ...string) {
	msgRet := "操作成功"
	data := make(map[string]interface{}) //初始化为空{}
	if len(msg) > 0 {
		for _, v := range msg {
			msgRet = ConcatStrings(v, " ")
		}
	}
	g.C.JSON(http.StatusOK, Response{
		Code: 200,
		Msg:  msgRet,
		Data: data,
	})
	return
}

// 填写成功的response信息，data不为空，code是200
func (g *Gin) ResponseSucMsg(data interface{}, msg ...string) {
	msgRet := "操作成功"

	if len(msg) > 0 {
		for _, v := range msg {
			msgRet = ConcatStrings(v, " ")
		}
	}
	g.C.JSON(http.StatusOK, Response{
		Code: 200,
		Msg:  msgRet,
		Data: data,
	})
	return
}

// 填写失败的response信息，code是500
func (g *Gin) ResponseErr(msg ...string) {
	msgRet := "服务端发生错误"
	lenParams := len(msg)
	data := make(map[string]interface{}) //初始化为空{}

	if lenParams > 0 {
		for _, v := range msg {
			msgRet = ConcatStrings(v, " ")
		}
	}

	g.C.JSON(http.StatusOK, Response{
		Code: 500,
		Msg:  msgRet,
		Data: data,
	})
	g.C.Abort() //提前结束后续的handler链条
	return
}

// 填写失败的response信息，code是500
func (g *Gin) ResponseErrMsg(data interface{}, msg ...string) {
	msgRet := "服务端发生错误"
	lenParams := len(msg)
	if lenParams > 0 {
		for _, v := range msg {
			msgRet = ConcatStrings(v, " ")
		}
	}

	g.C.JSON(http.StatusOK, Response{
		Code: 500,
		Msg:  msgRet,
		Data: data,
	})
	g.C.Abort() //提前结束后续的handler链条
	return
}

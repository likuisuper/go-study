package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type BaseController struct {
}

func (c *BaseController) ResponseSuccess(g *gin.Context) {
	g.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": "",
		"msg":  "操作成功",
	})
}

func (c *BaseController) ResponseData(g *gin.Context, data interface{}) {
	g.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": data,
		"msg":  "操作成功",
	})
}

func (c *BaseController) ResponseFailureForParameter(g *gin.Context, err interface{}) {
	g.JSON(http.StatusForbidden, gin.H{
		"code": 403,
		"data": "",
		"msg":  err,
	})
}

func (c *BaseController) ResponseFailureForFuncErr(g *gin.Context, err interface{}) {
	g.JSON(http.StatusInternalServerError, gin.H{
		"code": 500,
		"data": "",
		"msg":  err,
	})
}

func (c *BaseController) ResponseFailure(g *gin.Context, httpCode, code int, err interface{}) {
	g.JSON(httpCode, gin.H{
		"code": code,
		"data": "",
		"msg":  err,
	})
}

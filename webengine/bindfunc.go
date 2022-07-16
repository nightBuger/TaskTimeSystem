package webengine

import (
	gin "github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func Register(uri string, handler func(ctx *gin.Context), method string) {
	switch {
	case strings.ToUpper(method) == "GET":
		GinInstance.GET(uri, handler)
	case strings.ToUpper(method) == "POST":
		GinInstance.POST(uri, handler)
	case strings.ToUpper(method) == "PUT":
		GinInstance.PUT(uri, handler)
	case strings.ToUpper(method) == "DELETE":
		GinInstance.DELETE(uri, handler)
	default:
		Logger.Fatal("错误的method:", method, ",无法完成绑定")
	}
}

func JSONOK(c *gin.Context, any interface{}) {
	c.JSON(http.StatusOK, any)
}

func JSONBadRequest(c *gin.Context, err error) {
	c.JSON(http.StatusBadRequest, gin.H{
		"msg": err.Error(),
	})
}

func JSONSqlFailed(c *gin.Context, err error) {
	c.JSON(http.StatusInternalServerError, gin.H{
		"msg": err.Error(),
	})
}

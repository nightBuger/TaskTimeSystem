package service

import (
	"TaskTimeSystem/webengine"
	"github.com/gin-gonic/gin"
)

func init() {
	webengine.Register("/user/userinfo", GetUserInfo, "GET")
}

func GetUserInfo(c *gin.Context) {
	var json UserInfoGet

	if err := c.ShouldBindQuery(&json); err != nil {
		webengine.JSONBadRequest(c, err)
		return
	}
	var base SqlResult
	var res []UserInfoGetRes
	base.ResultSlice = &res

	webengine.JSONOK(c, base)
}

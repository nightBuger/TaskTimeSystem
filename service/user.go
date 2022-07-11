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
	var res UserInfoGetRes
	res.Id = json.Id
	res.Name = "原野"
	res.Level = 3
	res.Phonenum = "13301346616"
	webengine.JSONOK(c, res)
}

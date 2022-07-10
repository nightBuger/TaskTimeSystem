package service

import (
	"TaskTimeSystem/webengine"
	"github.com/gin-gonic/gin"
)

func init() {
	webengine.Register("/user/userinfo", GetUserInfo, "GET")
}

type UserInfoGet struct {
	Id string `json:"id" form:"id" binding:"required"`
}
type UserInfoGetRes struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Level    int    `json:"level"`
	Role     string `json:"role"`
	RoleId   string `json:"role_id"`
	Phonenum string `json:"phonenum"`
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
	res.Role = "系统管理员"
	res.RoleId = "1"
	res.Phonenum = "13301346616"
	webengine.JSONOK(c, res)
}

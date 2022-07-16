package service

import (
	"TaskTimeSystem/service/sqlcreator"
	"TaskTimeSystem/webengine"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
)

func init() {
	webengine.Register("/user/userinfo", GetUserInfo, "GET")
	webengine.Register("/user/userlist", GetUserList, "GET")
}

func GetUserInfo(c *gin.Context) {
	var json UserInfoGet

	if err := c.ShouldBindQuery(&json); err != nil {
		webengine.JSONBadRequest(c, err)
		return
	}
	var dr sqlcreator.Selecter
	dr.Table("t_userb").Select("*").Equal("userid", json.Id)
	sqlstr, err := dr.Create()
	if err != nil {
		webengine.JSONBadRequest(c, err)
		return
	}
	res, err := webengine.QueryRow(sqlstr, UserInfoGetRes{})
	if err != nil {
		webengine.JSONSqlFailed(c, err)
		return
	}

	if res.RowCount <= 0 {
		webengine.JSONSqlFailed(c, errors.New(fmt.Sprintf("无法找到用户[%s]", json.Id)))
		return
	}

	webengine.JSONOK(c, res)
}
func GetUserList(c *gin.Context) {
	var json UserListGet

	if err := c.ShouldBindQuery(&json); err != nil {
		webengine.JSONBadRequest(c, err)
		return
	}
	var dr sqlcreator.Selecter
	dr.Table("t_userb").Select("*")
	dr.LikeNotEmpty("userid", json.Id).LikeNotEmpty("username", json.Name)
	//dr.SetPage(json.PageNo, json.PageCount)
	sqlstr, err := dr.Create()
	if err != nil {
		webengine.JSONBadRequest(c, err)
		return
	}
	res, err := webengine.Query(sqlstr, UserListGetRes{})
	if err != nil {
		webengine.JSONSqlFailed(c, err)
		return
	}
	webengine.JSONOK(c, res)
}

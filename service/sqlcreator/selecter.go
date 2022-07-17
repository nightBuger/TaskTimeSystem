package sqlcreator

import (
	"TaskTimeSystem/webengine"
	"errors"
	"fmt"
)

type Selecter struct {
	likeList   map[string]string
	equalList  map[string]interface{}
	selectList []string
	tableName  string
	limit      [2]int
	bFoundRows bool
}

//基础设置
func (this *Selecter) Select(field string) *Selecter {
	this.selectList = append(this.selectList, field)
	return this
}
func (this *Selecter) Table(tableName string) *Selecter {
	this.tableName = tableName
	return this
}

//必选筛选项
func (this *Selecter) Like(field string, filter string) *Selecter {
	if this.likeList == nil {
		this.likeList = make(map[string]string)
	}
	this.likeList[field] = filter
	return this
}
func (this *Selecter) Equal(field string, filter interface{}) *Selecter {
	if this.equalList == nil {
		this.equalList = make(map[string]interface{})
	}
	this.equalList[field] = filter
	return this
}

//非空才筛选项
func (this *Selecter) LikeNotEmpty(field string, filter string) *Selecter {
	if this.likeList == nil {
		this.likeList = make(map[string]string)
	}
	if filter != "" {
		this.likeList[field] = filter
	}
	return this
}
func (this *Selecter) EqualNotEmpty(field string, filter string) *Selecter {
	if this.equalList == nil {
		this.equalList = make(map[string]interface{})
	}
	if filter != "" {
		this.equalList[field] = filter
	}
	return this
}

//分页设置
func (this *Selecter) SetPage(pageno, pagecount int) *Selecter {
	if pageno <= 0 || pagecount <= 0 {
		webengine.Logger.Error("Pageno aor Pagecount is not valid,please check code,", pageno, pagecount)
		this.bFoundRows = false
		return this
	}
	this.bFoundRows = true
	this.limit[1] = pagecount
	this.limit[0] = (pageno - 1) * pagecount
	return this
}

func (this *Selecter) Create() (sqlstr string, err error) {
	if this.tableName == "" || len(this.selectList) == 0 {
		return sqlstr, errors.New("非法的sql")
	}

	if this.bFoundRows {
		sqlstr = "select SQL_CALC_FOUND_ROWS"
	} else {
		sqlstr = "select"
	}

	comma := " "
	for _, s := range this.selectList {
		sqlstr += fmt.Sprintf("%s%s", comma, s)
		comma = ","
	}
	sqlstr += fmt.Sprintf(" from %s where 1=1 ", this.tableName)

	for field, filter := range this.equalList {
		switch filter.(type) {
		case string:
			sqlstr += fmt.Sprintf(" and %s='%s'", field, filter.(string))
		case int:
			sqlstr += fmt.Sprintf(" and %s=%d", field, filter.(int))
		default:
			webengine.Logger.Error(filter, " is not a vaild argument , please check code")
		}
	}

	for field, filter := range this.likeList {
		sqlstr += fmt.Sprintf(" and %s like '%%%s%%'", field, filter)
	}

	if this.bFoundRows {
		sqlstr += fmt.Sprintf(" limit %d,%d;SELECT FOUND_ROWS();", this.limit[0], this.limit[1])
	}
	webengine.Logger.Debug(sqlstr)
	return
}
